package detector

import "devdocgen/internal/model"

func DetectProject(root string) *model.ProjectMetadata {
    meta := &model.ProjectMetadata{
        RootPath: root,
    }

    DetectName(root, meta)
    DetectGo(root, meta)
    DetectNode(root, meta)
    DetectNodeScripts(root, meta)
    DetectEnv(root, meta)
    DetectDocker(root, meta)
    meta.License = DetectLicense(root)


    return meta
}
