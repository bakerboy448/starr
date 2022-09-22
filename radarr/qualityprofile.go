package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpQualityProfile = APIver + "/qualityProfile"

// QualityProfile is applied to Movies.
type QualityProfile struct {
	ID                int64               `json:"id"`
	Name              string              `json:"name"`
	UpgradeAllowed    bool                `json:"upgradeAllowed"`
	Cutoff            int64               `json:"cutoff"`
	Qualities         []*starr.Quality    `json:"items"`
	MinFormatScore    int64               `json:"minFormatScore"`
	CutoffFormatScore int64               `json:"cutoffFormatScore"`
	FormatItems       []*starr.FormatItem `json:"formatItems,omitempty"`
	Language          *starr.Value        `json:"language"`
}

// GetQualityProfiles returns all configured quality profiles.
func (r *Radarr) GetQualityProfiles() ([]*QualityProfile, error) {
	return r.GetQualityProfilesContext(context.Background())
}

// GetQualityProfilesContext returns all configured quality profiles.
func (r *Radarr) GetQualityProfilesContext(ctx context.Context) ([]*QualityProfile, error) {
	var output []*QualityProfile

	req := starr.Request{URI: bpQualityProfile}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// AddQualityProfile updates a quality profile in place.
func (r *Radarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	return r.AddQualityProfileContext(context.Background(), profile)
}

// AddQualityProfileContext updates a quality profile in place.
func (r *Radarr) AddQualityProfileContext(ctx context.Context, profile *QualityProfile) (int64, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return 0, fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	var output QualityProfile

	req := starr.Request{URI: bpQualityProfile, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return 0, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (r *Radarr) UpdateQualityProfile(profile *QualityProfile) error {
	return r.UpdateQualityProfileContext(context.Background(), profile)
}

// UpdateQualityProfileContext updates a quality profile in place.
func (r *Radarr) UpdateQualityProfileContext(ctx context.Context, profile *QualityProfile) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	var output interface{}

	req := starr.Request{URI: path.Join(bpQualityProfile, fmt.Sprint(profile.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return nil
}

// DeleteQualityProfile deletes a quality profile.
func (r *Radarr) DeleteQualityProfile(profileID int64) error {
	return r.DeleteQualityProfileContext(context.Background(), profileID)
}

// DeleteQualityProfileContext deletes a quality profile.
func (r *Radarr) DeleteQualityProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpQualityProfile, fmt.Sprint(profileID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
