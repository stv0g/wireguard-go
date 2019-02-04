// Code generated by 'go generate'; DO NOT EDIT.

package setupapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modsetupapi = windows.NewLazySystemDLL("setupapi.dll")

	procSetupDiCreateDeviceInfoListExW  = modsetupapi.NewProc("SetupDiCreateDeviceInfoListExW")
	procSetupDiGetDeviceInfoListDetailW = modsetupapi.NewProc("SetupDiGetDeviceInfoListDetailW")
	procSetupDiCreateDeviceInfoW        = modsetupapi.NewProc("SetupDiCreateDeviceInfoW")
	procSetupDiEnumDeviceInfo           = modsetupapi.NewProc("SetupDiEnumDeviceInfo")
	procSetupDiDestroyDeviceInfoList    = modsetupapi.NewProc("SetupDiDestroyDeviceInfoList")
	procSetupDiGetClassDevsExW          = modsetupapi.NewProc("SetupDiGetClassDevsExW")
	procSetupDiCallClassInstaller       = modsetupapi.NewProc("SetupDiCallClassInstaller")
	procSetupDiOpenDevRegKey            = modsetupapi.NewProc("SetupDiOpenDevRegKey")
	procSetupDiGetDeviceInstallParamsW  = modsetupapi.NewProc("SetupDiGetDeviceInstallParamsW")
	procSetupDiGetClassInstallParamsW   = modsetupapi.NewProc("SetupDiGetClassInstallParamsW")
	procSetupDiSetDeviceInstallParamsW  = modsetupapi.NewProc("SetupDiSetDeviceInstallParamsW")
	procSetupDiSetClassInstallParamsW   = modsetupapi.NewProc("SetupDiSetClassInstallParamsW")
	procSetupDiClassNameFromGuidExW     = modsetupapi.NewProc("SetupDiClassNameFromGuidExW")
	procSetupDiClassGuidsFromNameExW    = modsetupapi.NewProc("SetupDiClassGuidsFromNameExW")
	procSetupDiGetSelectedDevice        = modsetupapi.NewProc("SetupDiGetSelectedDevice")
	procSetupDiSetSelectedDevice        = modsetupapi.NewProc("SetupDiSetSelectedDevice")
)

func setupDiCreateDeviceInfoListEx(ClassGUID *windows.GUID, hwndParent uintptr, MachineName *uint16, Reserved uintptr) (handle DevInfo, err error) {
	r0, _, e1 := syscall.Syscall6(procSetupDiCreateDeviceInfoListExW.Addr(), 4, uintptr(unsafe.Pointer(ClassGUID)), uintptr(hwndParent), uintptr(unsafe.Pointer(MachineName)), uintptr(Reserved), 0, 0)
	handle = DevInfo(r0)
	if handle == DevInfo(windows.InvalidHandle) {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetDeviceInfoListDetail(DeviceInfoSet DevInfo, DeviceInfoSetDetailData *_SP_DEVINFO_LIST_DETAIL_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetDeviceInfoListDetailW.Addr(), 2, uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceInfoSetDetailData)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiCreateDeviceInfo(DeviceInfoSet DevInfo, DeviceName *uint16, ClassGUID *windows.GUID, DeviceDescription *uint16, hwndParent uintptr, CreationFlags DICD, DeviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall9(procSetupDiCreateDeviceInfoW.Addr(), 7, uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceName)), uintptr(unsafe.Pointer(ClassGUID)), uintptr(unsafe.Pointer(DeviceDescription)), uintptr(hwndParent), uintptr(CreationFlags), uintptr(unsafe.Pointer(DeviceInfoData)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiEnumDeviceInfo(DeviceInfoSet DevInfo, MemberIndex uint32, DeviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiEnumDeviceInfo.Addr(), 3, uintptr(DeviceInfoSet), uintptr(MemberIndex), uintptr(unsafe.Pointer(DeviceInfoData)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiDestroyDeviceInfoList(DeviceInfoSet DevInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiDestroyDeviceInfoList.Addr(), 1, uintptr(DeviceInfoSet), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetClassDevsEx(ClassGUID *windows.GUID, Enumerator *uint16, hwndParent uintptr, Flags DIGCF, DeviceInfoSet DevInfo, MachineName *uint16, reserved uintptr) (handle DevInfo, err error) {
	r0, _, e1 := syscall.Syscall9(procSetupDiGetClassDevsExW.Addr(), 7, uintptr(unsafe.Pointer(ClassGUID)), uintptr(unsafe.Pointer(Enumerator)), uintptr(hwndParent), uintptr(Flags), uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(MachineName)), uintptr(reserved), 0, 0)
	handle = DevInfo(r0)
	if handle == DevInfo(windows.InvalidHandle) {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiCallClassInstaller(InstallFunction DI_FUNCTION, DeviceInfoSet DevInfo, DeviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiCallClassInstaller.Addr(), 3, uintptr(InstallFunction), uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceInfoData)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiOpenDevRegKey(DeviceInfoSet DevInfo, DeviceInfoData *SP_DEVINFO_DATA, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (key windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procSetupDiOpenDevRegKey.Addr(), 6, uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceInfoData)), uintptr(Scope), uintptr(HwProfile), uintptr(KeyType), uintptr(samDesired))
	key = windows.Handle(r0)
	if key == windows.InvalidHandle {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetDeviceInstallParams(DeviceInfoSet DevInfo, DeviceInfoData *SP_DEVINFO_DATA, DeviceInstallParams *_SP_DEVINSTALL_PARAMS) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetDeviceInstallParamsW.Addr(), 3, uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceInfoData)), uintptr(unsafe.Pointer(DeviceInstallParams)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiGetClassInstallParams(DeviceInfoSet DevInfo, DeviceInfoData *SP_DEVINFO_DATA, ClassInstallParams *SP_CLASSINSTALL_HEADER, ClassInstallParamsSize uint32, RequiredSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiGetClassInstallParamsW.Addr(), 5, uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceInfoData)), uintptr(unsafe.Pointer(ClassInstallParams)), uintptr(ClassInstallParamsSize), uintptr(unsafe.Pointer(RequiredSize)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiSetDeviceInstallParams(DeviceInfoSet DevInfo, DeviceInfoData *SP_DEVINFO_DATA, DeviceInstallParams *_SP_DEVINSTALL_PARAMS) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiSetDeviceInstallParamsW.Addr(), 3, uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceInfoData)), uintptr(unsafe.Pointer(DeviceInstallParams)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiSetClassInstallParams(DeviceInfoSet DevInfo, DeviceInfoData *SP_DEVINFO_DATA, ClassInstallParams *SP_CLASSINSTALL_HEADER, ClassInstallParamsSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiSetClassInstallParamsW.Addr(), 4, uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceInfoData)), uintptr(unsafe.Pointer(ClassInstallParams)), uintptr(ClassInstallParamsSize), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiClassNameFromGuidEx(ClassGUID *windows.GUID, ClassName *uint16, ClassNameSize uint32, RequiredSize *uint32, MachineName *uint16, Reserved uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiClassNameFromGuidExW.Addr(), 6, uintptr(unsafe.Pointer(ClassGUID)), uintptr(unsafe.Pointer(ClassName)), uintptr(ClassNameSize), uintptr(unsafe.Pointer(RequiredSize)), uintptr(unsafe.Pointer(MachineName)), uintptr(Reserved))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiClassGuidsFromNameEx(ClassName *uint16, ClassGuidList *windows.GUID, ClassGuidListSize uint32, RequiredSize *uint32, MachineName *uint16, Reserved uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiClassGuidsFromNameExW.Addr(), 6, uintptr(unsafe.Pointer(ClassName)), uintptr(unsafe.Pointer(ClassGuidList)), uintptr(ClassGuidListSize), uintptr(unsafe.Pointer(RequiredSize)), uintptr(unsafe.Pointer(MachineName)), uintptr(Reserved))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetSelectedDevice(DeviceInfoSet DevInfo, DeviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetSelectedDevice.Addr(), 2, uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceInfoData)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiSetSelectedDevice(DeviceInfoSet DevInfo, DeviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiSetSelectedDevice.Addr(), 2, uintptr(DeviceInfoSet), uintptr(unsafe.Pointer(DeviceInfoData)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
