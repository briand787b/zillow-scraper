package model

import (
	"context"
	"fmt"
	"math"
	"zcrapr/core/perr"
	"zcrapr/core/plog"

	"github.com/pkg/errors"
)

// Property represents a Property
type Property struct {
	ID      string
	URL     string
	Acreage int
	Address string // TODO: make this more granular in model, but keep as string in Redis

	captures []Capture
}

// GetPropertyByID gets a Property by its ID
func GetPropertyByID(ctx context.Context, l plog.Logger, id string, ps PropertyStore) (*Property, error) {
	if id == "" {
		return nil, perr.NewErrInvalid("cannot search with an empty ID")
	}

	p, err := ps.GetPropertyByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not get property by ID from store")
	}

	return p, nil
}

// GetAllProperties retrieves all Properties
func GetAllProperties(ctx context.Context, l plog.Logger, take int, ps PropertyStore) ([]Property, error) {
	// if skip < 0 {
	// 	return nil, perr.NewErrInvalid("skip cannot be negative")
	// }

	if take < 1 {
		take = math.MaxInt32
	}

	propIDs, err := ps.GetAllPropertyIDs(ctx, take)
	if err != nil {
		return nil, errors.Wrap(err, "could not get all properties")
	}

	props := make([]Property, len(propIDs))
	for i, propID := range propIDs {
		prop, err := ps.GetPropertyByID(ctx, propID)
		if err != nil {
			return nil, errors.Wrap(err, "could not get property by ID")
		}

		props[i] = *prop
	}

	return props, nil
}

// GetPropertiesByAddress gets all properties that match the address
func GetPropertiesByAddress(ctx context.Context, l plog.Logger, address string, ps PropertyStore) ([]Property, error) {
	if address == "" {
		return nil, perr.NewErrInvalid("could not search by address with empty address")
	}

	props, err := ps.GetPropertiesByAddress(ctx, address)
	if err != nil {
		return nil, errors.Wrap(err, "could not get properties by address")
	}

	return props, nil
}

// LoadPropertyByFields loads a Property from the database using the fields of given Property.  This function is useful
// for loading Property records into memory when the caller is unsure if the Property reference it holds accurately
// represents a database record, or if it even exists.  It goes through the fields of the Property (in order of
// specificity) and searches on the first field that is non-empty.  If that field does not return a valid Property,
// this function returns an error
func LoadPropertyByFields(ctx context.Context, l plog.Logger, p *Property, ps PropertyStore) (*Property, error) {
	id := p.ID
	var err error
	if id == "" {
		if p.Address != "" {
			if id, err = ps.GetPropertyIDByAddress(ctx, p.Address); err != nil {
				return nil, errors.Wrap(err, "could not get Property ID by address")
			}
		} else if p.URL != "" {
			if id, err = ps.GetPropertyIDByURL(ctx, p.URL); err != nil {
				return nil, errors.Wrap(err, "could not get Property ID by URL")
			}
		} else {
			// there are no fields provided that can be used to load
			// a Property and this function will return an error
			return nil, perr.NewErrInvalid("provided Property has no uniquely identifiable fields on it")
		}
	}

	loaded, err := ps.GetPropertyByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "could not load Property by ID")
	}

	return loaded, nil
}

// Save saves a property to the database
func (p *Property) Save(ctx context.Context, l plog.Logger, ps PropertyStore) error {
	if p.URL == "" {
		return perr.NewErrInvalid("URL cannot be empty")
	}

	if p.Acreage < 1 {
		return perr.NewErrInvalid("properties must have at least one acre (rounded up)")
	}

	if p.Address == "" {
		return perr.NewErrInvalid("Address cannot be empty string")
	}

	if p.ID == "" {
		return p.insert(ctx, l, ps)
	}

	return p.update(ctx, ps)
}

// GetCaptures retrieves all the already loaded Captures
func (p *Property) GetCaptures() []Capture {
	return p.captures
}

// LoadCaptures loads all captures into the Property receiever
func (p *Property) LoadCaptures(ctx context.Context, l plog.Logger, cs CaptureStore) error {
	caps, err := cs.GetAllCapturesByPropertyID(ctx, p.ID)
	if err != nil {
		return errors.Wrap(err, "could not get Captures")
	}

	p.captures = caps
	return nil
}

// helper methods

func (p *Property) insert(ctx context.Context, l plog.Logger, ps PropertyStore) error {
	_, err := ps.GetPropertyIDByURL(ctx, p.URL)
	if err != nil && !perr.SameType(err, perr.ErrNotFound) {
		return errors.Wrap(err, "could not get property ID from URL")
	} else if err == nil {
		return perr.NewErrInvalid(fmt.Sprintf("URL %s already exists in the database", p.URL))
	}

	_, err = ps.GetPropertyIDByAddress(ctx, p.Address)
	if err != nil && !perr.SameType(err, perr.ErrNotFound) {
		return errors.Wrap(err, "could not get property ID from address")
	} else if err == nil {
		return perr.NewErrInvalid(fmt.Sprintf("address %s already exists in the database", p.Address))
	}

	if err := ps.InsertProperty(ctx, p); err != nil {
		return errors.Wrap(err, "could not insert Property")
	}

	return nil
}

func (p *Property) update(ctx context.Context, ps PropertyStore) error {
	oldProp, err := ps.GetPropertyByID(ctx, p.ID)
	if err != nil {
		return errors.Wrap(err, "could not get old property by ID")
	}

	if oldProp.Acreage != p.Acreage {
		return perr.NewErrInvalid("acreage cannot be mutated without invalidating existing captures")
	}

	if oldProp.Address != p.Address {
		return perr.NewErrInvalid("address cannot be mutated without invalidating existing captures")
	}

	if err := ps.UpdateProperty(ctx, p); err != nil {
		return errors.Wrap(err, "could not update Property")
	}

	return nil
}
