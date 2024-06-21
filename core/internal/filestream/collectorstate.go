package filestream

import "github.com/wandb/wandb/core/internal/sparselist"

// CollectorState is the filestream's buffered data.
type CollectorState struct {
	HistoryLineNum int      // Line number where to append run history.
	HistoryLines   []string // Lines to append to run history.

	EventsLineNum int      // Line number where to append run system metrics.
	EventsLines   []string // Lines to append to run system metrics.

	// Lines to update in the run's console logs file.
	ConsoleLogUpdates sparselist.SparseList[string]

	// Offset to add to all console output line numbers.
	//
	// This is used for resumed runs, where we want to append to the original
	// logs.
	ConsoleLogLineOffset int

	SummaryLineNum int    // Line number where to write the run summary.
	LatestSummary  string // The run's updated summary, or the empty string.

	// UploadedFiles are files for which uploads have finished.
	UploadedFiles []string

	HasPreempting bool
	Preempting    bool

	// ExitCode is the run's script's exit code if any.
	//
	// This is sent with the final transmission.
	ExitCode *int32

	// Complete is the run's script's completion status if any.
	//
	// This is sent with the final transmission.
	Complete *bool
}

func NewCollectorState(initialOffsets FileStreamOffsetMap) CollectorState {
	state := CollectorState{}

	if initialOffsets != nil {
		state.HistoryLineNum = initialOffsets[HistoryChunk]
		state.EventsLineNum = initialOffsets[EventsChunk]
		state.ConsoleLogLineOffset = initialOffsets[OutputChunk]
		state.SummaryLineNum = initialOffsets[SummaryChunk]
	}

	return state
}

// CollectorStateUpdate is a mutation to a CollectorState.
type CollectorStateUpdate interface {
	// Apply modifies the collector state.
	Apply(*CollectorState)
}

// MakeRequest moves buffered data into an API request and returns it.
//
// Returns a boolean that's true if the request is non-empty.
func (s *CollectorState) MakeRequest(isDone bool) (*FsTransmitData, bool) {
	files := make(map[string]FsTransmitFileData)
	addLines := func(chunkType ChunkTypeEnum, lineNum int, lines []string) {
		if len(lines) == 0 {
			return
		}
		files[chunkFilename[chunkType]] = FsTransmitFileData{
			Offset:  lineNum,
			Content: lines,
		}
	}

	addLines(HistoryChunk, s.HistoryLineNum, s.HistoryLines)
	s.HistoryLineNum += len(s.HistoryLines)
	s.HistoryLines = nil

	addLines(EventsChunk, s.EventsLineNum, s.EventsLines)
	s.EventsLineNum += len(s.EventsLines)
	s.EventsLines = nil

	if s.ConsoleLogUpdates.Len() > 0 {
		// We can only upload one run of lines at a time, unfortunately.
		run := s.ConsoleLogUpdates.ToRuns()[0]
		files[chunkFilename[OutputChunk]] = FsTransmitFileData{
			Offset:  run.Start + s.ConsoleLogLineOffset,
			Content: run.Items,
		}

		for i := run.Start; i < run.Start+len(run.Items); i++ {
			s.ConsoleLogUpdates.Delete(i)
		}
	}

	if s.LatestSummary != "" {
		// We always write to the same line in the summary file.
		//
		// The reason this isn't always line 0 is for compatibility with old
		// runs where we appended to the summary file. In that case, we want
		// to update the last line, since all other lines are ignored. This
		// applies to resumed runs.
		addLines(SummaryChunk, s.SummaryLineNum, []string{s.LatestSummary})
		s.LatestSummary = ""
	}

	transmitData := FsTransmitData{}
	hasData := false

	if len(files) > 0 {
		transmitData.Files = files
		hasData = true
	}

	if len(s.UploadedFiles) > 0 {
		transmitData.Uploaded = s.UploadedFiles
		s.UploadedFiles = nil
		hasData = true
	}

	if s.HasPreempting {
		transmitData.Preempting = s.Preempting
		hasData = true
	}

	if isDone {
		transmitData.Exitcode = s.ExitCode
		transmitData.Complete = s.Complete
		hasData = true
	}

	return &transmitData, hasData
}
