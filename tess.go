package tess

/*
#cgo CFLAGS: -I${SRCDIR}/etc/include
#cgo CXXFLAGS: -I${SRCDIR}/etc/include
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/etc/lib/linux_amd64 -lyogacore -lstdc++ -lm
#cgo linux,arm64 LDFLAGS: -L${SRCDIR}/etc/lib/linux_arm64 -lyogacore -lstdc++ -lm
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/etc/lib/darwin_amd64 -lyogacore -lc++
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/etc/lib/darwin_arm64 -lyogacore -lc++
#cgo windows,amd64 LDFLAGS: ${SRCDIR}/etc/lib/windows_amd64/yogacore.lib
#cgo windows,arm64 LDFLAGS: ${SRCDIR}/etc/lib/windows_arm64/yogacore.lib
#include <yoga/Yoga.h>
*/
import "C"
