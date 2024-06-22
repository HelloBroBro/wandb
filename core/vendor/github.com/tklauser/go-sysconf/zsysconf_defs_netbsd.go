// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs sysconf_defs_netbsd.go

//go:build netbsd

package sysconf

const (
	SC_ARG_MAX                      = 0x1
	SC_CHILD_MAX                    = 0x2
	SC_NGROUPS_MAX                  = 0x4
	SC_OPEN_MAX                     = 0x5
	SC_JOB_CONTROL                  = 0x6
	SC_SAVED_IDS                    = 0x7
	SC_VERSION                      = 0x8
	SC_BC_BASE_MAX                  = 0x9
	SC_BC_DIM_MAX                   = 0xa
	SC_BC_SCALE_MAX                 = 0xb
	SC_BC_STRING_MAX                = 0xc
	SC_COLL_WEIGHTS_MAX             = 0xd
	SC_EXPR_NEST_MAX                = 0xe
	SC_LINE_MAX                     = 0xf
	SC_RE_DUP_MAX                   = 0x10
	SC_2_VERSION                    = 0x11
	SC_2_C_BIND                     = 0x12
	SC_2_C_DEV                      = 0x13
	SC_2_CHAR_TERM                  = 0x14
	SC_2_FORT_DEV                   = 0x15
	SC_2_FORT_RUN                   = 0x16
	SC_2_LOCALEDEF                  = 0x17
	SC_2_SW_DEV                     = 0x18
	SC_2_UPE                        = 0x19
	SC_STREAM_MAX                   = 0x1a
	SC_TZNAME_MAX                   = 0x1b
	SC_PAGESIZE                     = 0x1c
	SC_PAGE_SIZE                    = 0x1c
	SC_FSYNC                        = 0x1d
	SC_XOPEN_SHM                    = 0x1e
	SC_SYNCHRONIZED_IO              = 0x1f
	SC_IOV_MAX                      = 0x20
	SC_MAPPED_FILES                 = 0x21
	SC_MEMLOCK                      = 0x22
	SC_MEMLOCK_RANGE                = 0x23
	SC_MEMORY_PROTECTION            = 0x24
	SC_LOGIN_NAME_MAX               = 0x25
	SC_MONOTONIC_CLOCK              = 0x26
	SC_CLK_TCK                      = 0x27
	SC_ATEXIT_MAX                   = 0x28
	SC_THREADS                      = 0x29
	SC_SEMAPHORES                   = 0x2a
	SC_BARRIERS                     = 0x2b
	SC_TIMERS                       = 0x2c
	SC_SPIN_LOCKS                   = 0x2d
	SC_READER_WRITER_LOCKS          = 0x2e
	SC_GETGR_R_SIZE_MAX             = 0x2f
	SC_GETPW_R_SIZE_MAX             = 0x30
	SC_CLOCK_SELECTION              = 0x31
	SC_ASYNCHRONOUS_IO              = 0x32
	SC_AIO_LISTIO_MAX               = 0x33
	SC_AIO_MAX                      = 0x34
	SC_MESSAGE_PASSING              = 0x35
	SC_MQ_OPEN_MAX                  = 0x36
	SC_MQ_PRIO_MAX                  = 0x37
	SC_PRIORITY_SCHEDULING          = 0x38
	SC_THREAD_DESTRUCTOR_ITERATIONS = 0x39
	SC_THREAD_KEYS_MAX              = 0x3a
	SC_THREAD_STACK_MIN             = 0x3b
	SC_THREAD_THREADS_MAX           = 0x3c
	SC_THREAD_ATTR_STACKADDR        = 0x3d
	SC_THREAD_ATTR_STACKSIZE        = 0x3e
	SC_THREAD_PRIORITY_SCHEDULING   = 0x3f
	SC_THREAD_PRIO_INHERIT          = 0x40
	SC_THREAD_PRIO_PROTECT          = 0x41
	SC_THREAD_PROCESS_SHARED        = 0x42
	SC_THREAD_SAFE_FUNCTIONS        = 0x43
	SC_TTY_NAME_MAX                 = 0x44
	SC_HOST_NAME_MAX                = 0x45
	SC_PASS_MAX                     = 0x46
	SC_REGEXP                       = 0x47
	SC_SHELL                        = 0x48
	SC_SYMLOOP_MAX                  = 0x49

	SC_V6_ILP32_OFF32   = 0x4a
	SC_V6_ILP32_OFFBIG  = 0x4b
	SC_V6_LP64_OFF64    = 0x4c
	SC_V6_LPBIG_OFFBIG  = 0x4d
	SC_2_PBS            = 0x50
	SC_2_PBS_ACCOUNTING = 0x51
	SC_2_PBS_CHECKPOINT = 0x52
	SC_2_PBS_LOCATE     = 0x53
	SC_2_PBS_MESSAGE    = 0x54
	SC_2_PBS_TRACK      = 0x55

	SC_SPAWN                 = 0x56
	SC_SHARED_MEMORY_OBJECTS = 0x57

	SC_TIMER_MAX        = 0x58
	SC_SEM_NSEMS_MAX    = 0x59
	SC_CPUTIME          = 0x5a
	SC_THREAD_CPUTIME   = 0x5b
	SC_DELAYTIMER_MAX   = 0x5c
	SC_SIGQUEUE_MAX     = 0x5d
	SC_REALTIME_SIGNALS = 0x5e

	SC_PHYS_PAGES = 0x79

	SC_NPROCESSORS_CONF = 0x3e9
	SC_NPROCESSORS_ONLN = 0x3ea

	SC_SCHED_RT_TS   = 0x7d1
	SC_SCHED_PRI_MIN = 0x7d2
	SC_SCHED_PRI_MAX = 0x7d3
)

const (
	_MAXHOSTNAMELEN = 0x100
	_MAXLOGNAME     = 0x10
	_MAXSYMLINKS    = 0x20

	_POSIX_ARG_MAX                      = 0x1000
	_POSIX_CHILD_MAX                    = 0x19
	_POSIX_CPUTIME                      = 0x30db0
	_POSIX_DELAYTIMER_MAX               = 0x20
	_POSIX_PRIORITY_SCHEDULING          = 0x30db0
	_POSIX_REGEXP                       = 0x1
	_POSIX_SHARED_MEMORY_OBJECTS        = 0x0
	_POSIX_SHELL                        = 0x1
	_POSIX_SIGQUEUE_MAX                 = 0x20
	_POSIX_SPAWN                        = 0x31069
	_POSIX_THREAD_ATTR_STACKADDR        = 0x30db0
	_POSIX_THREAD_ATTR_STACKSIZE        = 0x30db0
	_POSIX_THREAD_CPUTIME               = 0x30db0
	_POSIX_THREAD_DESTRUCTOR_ITERATIONS = 0x4
	_POSIX_THREAD_KEYS_MAX              = 0x80
	_POSIX_THREAD_PRIO_PROTECT          = 0x30db0
	_POSIX_THREAD_SAFE_FUNCTIONS        = 0x30db0
	_POSIX_TIMER_MAX                    = 0x20
	_POSIX_VERSION                      = 0x30db0

	_POSIX2_VERSION = 0x30db0

	_FOPEN_MAX  = 0x14
	_NAME_MAX   = 0x1ff
	_RE_DUP_MAX = 0xff

	_BC_BASE_MAX      = 0x7fffffff
	_BC_DIM_MAX       = 0xffff
	_BC_SCALE_MAX     = 0x7fffffff
	_BC_STRING_MAX    = 0x7fffffff
	_COLL_WEIGHTS_MAX = 0x2
	_EXPR_NEST_MAX    = 0x20
	_LINE_MAX         = 0x800

	_GETGR_R_SIZE_MAX = 0x400
	_GETPW_R_SIZE_MAX = 0x400

	_PATH_DEV      = "/dev/"
	_PATH_ZONEINFO = "/usr/share/zoneinfo"

	_PASSWORD_LEN = 0x80
)

const _PC_NAME_MAX = 0x4
