package tess

/*
#cgo CFLAGS: -I${SRCDIR}/etc/include -I${SRCDIR}
#cgo CXXFLAGS: -I${SRCDIR}/etc/include -I${SRCDIR} -std=c++20
#cgo linux,amd64 LDFLAGS: -L${SRCDIR}/etc/lib/linux_amd64 -lyogacore -lstdc++ -lm
#cgo linux,arm64 LDFLAGS: -L${SRCDIR}/etc/lib/linux_arm64 -lyogacore -lstdc++ -lm
#cgo darwin,amd64 LDFLAGS: -L${SRCDIR}/etc/lib/darwin_amd64 -lyogacore -lc++
#cgo darwin,arm64 LDFLAGS: -L${SRCDIR}/etc/lib/darwin_arm64 -lyogacore -lc++
#cgo windows,amd64 LDFLAGS: -L${SRCDIR}/etc/lib/windows_amd64 -lyogacore -lstdc++ -lm
#cgo windows,arm64 LDFLAGS: -L${SRCDIR}/etc/lib/windows_arm64 -lyogacore -lstdc++ -lm
#include <yoga/Yoga.h>
#include "tess_ext.h"
*/
import "C"
