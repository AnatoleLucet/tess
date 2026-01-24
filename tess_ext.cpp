// Extension functions for tess that expose internal Yoga functionality

#include <yoga/Yoga.h>
#include <yoga/node/Node.h>

extern "C" {

void YGNodeSetDirtyExt(YGNodeRef nodeRef, bool isDirty) {
    static_cast<facebook::yoga::Node*>(nodeRef)->setDirty(isDirty);
}

}
