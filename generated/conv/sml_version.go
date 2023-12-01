// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package conv

import (
	conversion "github.com/satisfactorymodding/smr-api/conversion"
	generated "github.com/satisfactorymodding/smr-api/generated"
	ent "github.com/satisfactorymodding/smr-api/generated/ent"
)

type SMLVersionImpl struct{}

func (c *SMLVersionImpl) Convert(source *ent.SmlVersion) *generated.SMLVersion {
	var pGeneratedSMLVersion *generated.SMLVersion
	if source != nil {
		var generatedSMLVersion generated.SMLVersion
		generatedSMLVersion.ID = (*source).ID
		generatedSMLVersion.Version = (*source).Version
		generatedSMLVersion.SatisfactoryVersion = (*source).SatisfactoryVersion
		generatedSMLVersion.Stability = generated.VersionStabilities((*source).Stability)
		generatedSMLVersion.Link = (*source).Link
		var pGeneratedSMLVersionTargetList []*generated.SMLVersionTarget
		if (*source).Edges.Targets != nil {
			pGeneratedSMLVersionTargetList = make([]*generated.SMLVersionTarget, len((*source).Edges.Targets))
			for i := 0; i < len((*source).Edges.Targets); i++ {
				pGeneratedSMLVersionTargetList[i] = c.pEntSmlVersionTargetToPGeneratedSMLVersionTarget((*source).Edges.Targets[i])
			}
		}
		generatedSMLVersion.Targets = pGeneratedSMLVersionTargetList
		generatedSMLVersion.Changelog = (*source).Changelog
		generatedSMLVersion.Date = conversion.TimeToString((*source).Date)
		pString := (*source).BootstrapVersion
		generatedSMLVersion.BootstrapVersion = &pString
		generatedSMLVersion.EngineVersion = (*source).EngineVersion
		generatedSMLVersion.UpdatedAt = conversion.TimeToString((*source).UpdatedAt)
		generatedSMLVersion.CreatedAt = conversion.TimeToString((*source).CreatedAt)
		pGeneratedSMLVersion = &generatedSMLVersion
	}
	return pGeneratedSMLVersion
}
func (c *SMLVersionImpl) ConvertSlice(source []*ent.SmlVersion) []*generated.SMLVersion {
	var pGeneratedSMLVersionList []*generated.SMLVersion
	if source != nil {
		pGeneratedSMLVersionList = make([]*generated.SMLVersion, len(source))
		for i := 0; i < len(source); i++ {
			pGeneratedSMLVersionList[i] = c.Convert(source[i])
		}
	}
	return pGeneratedSMLVersionList
}
func (c *SMLVersionImpl) pEntSmlVersionTargetToPGeneratedSMLVersionTarget(source *ent.SmlVersionTarget) *generated.SMLVersionTarget {
	var pGeneratedSMLVersionTarget *generated.SMLVersionTarget
	if source != nil {
		var generatedSMLVersionTarget generated.SMLVersionTarget
		generatedSMLVersionTarget.VersionID = (*source).VersionID
		generatedSMLVersionTarget.TargetName = generated.TargetName((*source).TargetName)
		generatedSMLVersionTarget.Link = (*source).Link
		pGeneratedSMLVersionTarget = &generatedSMLVersionTarget
	}
	return pGeneratedSMLVersionTarget
}
