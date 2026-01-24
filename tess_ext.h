// Extension functions for tess that expose internal Yoga functionality

#ifndef TESS_EXT_H
#define TESS_EXT_H

#include <yoga/Yoga.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

void YGNodeSetDirtyExt(YGNodeRef nodeRef, bool isDirty);

#ifdef __cplusplus
}
#endif

#endif // TESS_EXT_H
