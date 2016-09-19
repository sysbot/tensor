package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
)

const DBC_INVENTORY_SOURCES = "inventory_sources"

// Inventory is the model for
// Inventory collection
// TODO: not implemented
type InventorySource struct {
	ID                 bson.ObjectId `bson:"_id" json:"id"`

	Source             string        `bson:"source" json:"source"`
	SourcePath         string        `bson:"source_path" json:"source_path"`
	SourceVars         string        `bson:"source_vars" json:"source_regions"`
	SourceRegions      string        `bson:"source_regions" json:"has_active_failures"`
	InstanceFilters    string        `bson:"instance_filters" json:"instance_filters"`
	GroupBy            string        `bson:"group_by" json:"group_by"`
	Overwrite          bool          `bson:"overwrite" json:"overwrite"`
	OverwriteVars      bool          `bson:"overwrite_vars" json:"overwrite_vars"`
	UpdateOnLaunch     bool          `bson:"update_on_launch" json:"update_on_launch"`
	UpdateCacheTimeout uint32        `bson:"update_cache_timeout" json:"update_cache_timeout"`
	CredentialID       bson.ObjectId `bson:"credential_id" json:"credential"`
	GroupID            bson.ObjectId `bson:"group_id" json:"group"`
	InventoryID        bson.ObjectId `bson:"inventory_id" json:"inventory"`
	SourceScriptID     bson.ObjectId `bson:"source_script_id" json:"source_script"`

	Type               string        `bson:"-" json:"type"`
	Url                string        `bson:"-" json:"url"`
	Related            gin.H         `bson:"-" json:"related"`
	SummaryFields      gin.H         `bson:"-" json:"summary_fields"`
}


func (is InventorySource) CreateIndexes()  {

}