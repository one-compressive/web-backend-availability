package repository

import (
	"fmt"
)

type Repository struct {
}

func NewRepository() (*Repository, error) {
	return &Repository{}, nil
}

type SystemImpact string

const (
	ImpactLow    SystemImpact = "Low"
	ImpactMedium SystemImpact = "Medium"
	ImpactHigh   SystemImpact = "High"
)

type ConfigType string

const (
	ConfigTypeSingle      ConfigType = "Single Node"
	ConfigTypeClustering  ConfigType = "Clustering"
	ConfigTypeReplication ConfigType = "Replication"
)

type Status string

const (
	StatusPublished Status = "Published"
	StatusDraft     Status = "Draft"
	StatusDeleted   Status = "Deleted"
)

type Component struct {
	ID               int
	Name             string
	Description      string
	ShortDescription string
	ConfigType       ConfigType
	UptimePercent    float32
	SystemImpact     SystemImpact
	Status           Status
	ImageURL         string
	VideoURL         string
	Likes            []int
}

func userIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i + 1
	}
	return ids
}

func (r *Repository) GetComponent(id int) (*Component, error) {
	components, err := r.GetComponents()
	if err != nil {
		return nil, err
	}

	for i := range components {
		if components[i].ID == id {
			return &components[i], nil
		}
	}

	return nil, fmt.Errorf("компонент не найден")
}

func (r *Repository) GetDraftComponent() (*Component, error) {
	components, err := r.GetComponents()
	if err != nil {
		return nil, err
	}

	for i := range components {
		if components[i].Status == StatusDraft {
			return &components[i], nil
		}
	}

	return nil, fmt.Errorf("черновик не найден")
}

func (r *Repository) GetComponents() ([]Component, error) {
	// имитируем работу с БД. Типа мы выполнили sql запрос и получили эти строки из БД
	components := []Component{
		{
			ID:               1,
			Name:             "High-Availability Database Cluster",
			Description:      "Multi-node active-active clustering for zero downtime during maintenance and automatic failover under node failure.",
			ShortDescription: "Multi-node active-active clustering architecture to ensure zero...",
			ConfigType:       ConfigTypeClustering,
			UptimePercent:    99.99,
			SystemImpact:     ImpactHigh,
			Status:           StatusPublished,
			ImageURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCRiVEMCVCQiVEMCVCOCVEMSU4MiVEMCVCQSVEMCVCMDEuanBnP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTE1MzQ0WiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTc2OGIxOGFiMzgyZTVhYTUyMmNmNWZmYWVjMWRiZmMzM2EyZTM5NDM4NTM0NjViNmFmOTNkYzAyYmM2NTNhMmQ",
			VideoURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCMiVEMCVCOCVEMCVCNCVEMCVCNSVEMCVCRTEubXA0P1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTEwNTEwWiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPWY3MzI0NWEzZDMwOTRmYTY2ZGYwMWQzMmYzNGRiM2NkY2ZiOTQwM2ExZmQxODdjN2E4YzA4NDFhN2MwY2QwNzg",
			Likes:            userIDs(12),
		},
		{
			ID:               2,
			Name:             "Web Server Node Alpha",
			Description:      "Single-node edge web tier used as a baseline for comparing availability gains after replication.",
			ShortDescription: "Single-node edge web tier baseline for availability comparison...",
			ConfigType:       ConfigTypeSingle,
			UptimePercent:    99.99,
			SystemImpact:     ImpactMedium,
			Status:           StatusPublished,
			ImageURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCRiVEMCVCQiVEMCVCOCVEMSU4MiVEMCVCQSVEMCVCMDIuanBnP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTE1NDA5WiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPWI2MWE5NDVhY2QwMjQ5MjQ0ZmI3YWYwMzA2NzlkMDg2MjYwOWQ0MmNhNjY0ZDQxMDEwNjY2NTdlOWJmYzY5MWY",
			VideoURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCMiVEMCVCOCVEMCVCNCVEMCVCNSVEMCVCRSUyMDMubXA0P1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTExODE1WiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTBkMGUxZjRmOTEyNDZmY2JkNDhhNGVlNGQzZTAxN2Y0ODk4OGQ2NGI2Y2IwMzE1ZWI4YmUwYzcxZTAyNGY1YzQ",
			Likes:            userIDs(1200),
		},
		{
			ID:               3,
			Name:             "Global Load Balancer",
			Description:      "Anycast load balancer across regional replicas to keep traffic reachable during partial outages.",
			ShortDescription: "Anycast load balancer across regional replicas...",
			ConfigType:       ConfigTypeReplication,
			UptimePercent:    99.95,
			SystemImpact:     ImpactHigh,
			Status:           StatusPublished,
			ImageURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCRiVEMCVCQiVEMCVCOCVEMSU4MiVEMCVCQSVEMCVCMDMuanBnP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTE1NDI5WiZYLUFtei1FeHBpcmVzPTQzMTk5JlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTYxNjQ0YTNhMmYzZjc5MDZkZDU4ZTlmYWU2NjUxYjk5OTZiZjU0MjlmMjliYThiNDg2YzJjNWU1MzFmZmM0NTY",
			VideoURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCMiVEMCVCOCVEMCVCNCVEMCVCNSVEMCVCRSUyMDMubXA0P1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTEzMjM2WiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTAzY2U3YzE3Y2YwYWFkZDVlYzk5ZTBhZTcxOGNiYmFlZmFkOGM1NTgwMDI1MTAzODNkOGUzNjVkYWQ3ZjNhODE",
			Likes:            userIDs(890),
		},
		{
			ID:               4,
			Name:             "Database Primary Replica",
			Description:      "Synchronous primary-replica pair for near-zero RPO and controlled failover drills.",
			ShortDescription: "Synchronous primary-replica pair for near-zero RPO...",
			ConfigType:       ConfigTypeReplication,
			UptimePercent:    99.91,
			SystemImpact:     ImpactHigh,
			Status:           StatusPublished,
			ImageURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCRiVEMCVCQiVEMCVCOCVEMSU4MiVEMCVCQSVEMCVCMDQuanBnP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTE1NDQ3WiZYLUFtei1FeHBpcmVzPTQzMTk5JlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTc3MjQ1ODE5MGIzMjMwOWEwZTEyNTFmMDRmNGE2MTJlNzRjNDdkOWEzNzgzNTZlMzFkZGNjNzc0YmNiOTg1ZTI",
			VideoURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCMiVEMCVCOCVEMCVCNCVEMCVCNSVEMCVCRSUyMDQubXA0P1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTEzMjU1WiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTZkZWU2MTYyN2Y1OTk1NGRlYWNjNjU1NmJmMTk2ZTAxZmNiZDQ2NjkxZDE4YjJmY2RhYTMyMmZiODg2YTljZWE",
			Likes:            userIDs(450),
		},
		{
			ID:               5,
			Name:             "Edge Caching Layer",
			Description:      "CDN caching tier that reduces origin load and masks short origin outages.",
			ShortDescription: "CDN caching tier masking short origin outages...",
			ConfigType:       ConfigTypeClustering,
			UptimePercent:    99.88,
			SystemImpact:     ImpactMedium,
			Status:           StatusPublished,
			ImageURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCRiVEMCVCQiVEMCVCOCVEMSU4MiVEMCVCQSVEMCVCMDUuanBnP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTE1NTAzWiZYLUFtei1FeHBpcmVzPTQzMTk5JlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTdjYTIxNjNmMTM0NTc4MDU1ZTUyZjhkMDk1YjI1ODhkNjZjMTBmN2I1OWZmZjA3ZjBiMDVhZDBlZjNlNzZhYWM",
			VideoURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCMiVEMCVCOCVEMCVCNCVEMCVCNSVEMCVCRSUyMDUubXA0P1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTEzMzE2WiZYLUFtei1FeHBpcmVzPTQzMTk5JlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPWY1MTQ1YjIzMDdmYWU1ZGE3MjUyYWMxYjczNzYwMTYxNGZmNTBkNGRmNDA5ZjA4NzMzZjI1Y2QyM2U2YjFmN2I",
			Likes:            userIDs(312),
		},
		{
			ID:               6,
			Name:             "Monitoring Probe Agent",
			Description:      "Draft single-node probe agent prepared for the next availability assessment scenario.",
			ShortDescription: "Draft probe agent for the next assessment scenario...",
			ConfigType:       ConfigTypeSingle,
			UptimePercent:    99.50,
			SystemImpact:     ImpactLow,
			Status:           StatusDraft,
			ImageURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCRiVEMCVCQiVEMCVCOCVEMSU4MiVEMCVCQSVEMCVCMDYuanBnP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTE1NTIzWiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTY2M2NiZDZjNDhhNmZhOTVkOGI4ZTlhZGQ4NDhkNjk1NzgwZTlmODllNDgwOTYxNTNkZTM5MTNkN2Y0NWZkMGI",
			VideoURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCMiVEMCVCOCVEMCVCNCVEMCVCNSVEMCVCRSUyMDYubXA0P1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTEzMzM2WiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTVkYTNmNzI0YzY1NDk0YjQ5NGY4YWZmY2M3Zjc3ZTE4ZTViM2RiY2JmNGI2NDBiOGQ4ZjNlNGIzOGM1MTE0MmY",
			Likes:            []int{},
		},
		{
			ID:               7,
			Name:             "Legacy Message Broker",
			Description:      "Deprecated broker kept only for history; must not appear in UI.",
			ShortDescription: "Deprecated broker removed from UI...",
			ConfigType:       ConfigTypeSingle,
			UptimePercent:    97.10,
			SystemImpact:     ImpactLow,
			Status:           StatusDeleted,
			ImageURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCRiVEMCVCQiVEMCVCOCVEMSU4MiVEMCVCQSVEMCVCMDcuanBnP1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTE1NTM5WiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTQzOWFmNjQ2NTU5MGI0Y2FhNzQ3YjMxZWY3OWViMTc4NTdhNjQ5NjM5MTk1NmI1MDBlOWRkZDI4NjM4ZDQwNDE",
			VideoURL:         "http://localhost:9001/api/v1/download-shared-object/aHR0cDovLzEyNy4wLjAuMTo5MDAwL21lZGlhLyVEMCVCMiVEMCVCOCVEMCVCNCVEMCVCNSVEMCVCRSUyMDYubXA0P1gtQW16LUFsZ29yaXRobT1BV1M0LUhNQUMtU0hBMjU2JlgtQW16LUNyZWRlbnRpYWw9R0NRM0xZWlVLNjRPWFA2RlZGMUslMkYyMDI2MDkwMiUyRnVzLWVhc3QtMSUyRnMzJTJGYXdzNF9yZXF1ZXN0JlgtQW16LURhdGU9MjAyNjA5MDJUMTEzMzM2WiZYLUFtei1FeHBpcmVzPTQzMjAwJlgtQW16LVNlY3VyaXR5LVRva2VuPWV5SmhiR2NpT2lKSVV6VXhNaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpoWTJObGMzTkxaWGtpT2lKSFExRXpURmxhVlVzMk5FOVlVRFpHVmtZeFN5SXNJbVY0Y0NJNk1UYzRPRE00TmpZeE9Dd2ljR0Z5Wlc1MElqb2ljbTl2ZENKOS5xT3NNT3J6ZFpkU1BVQUFnQ2NVcHY1My1mbTZqRGxBV1JiTTdnd0NzRnBFdC1DMXhUT01FVmNDdzFjdmo0MlVrUnBCMWtYM3g5aGIxSUw1VHA2M1E4USZYLUFtei1TaWduZWRIZWFkZXJzPWhvc3QmdmVyc2lvbklkPW51bGwmWC1BbXotU2lnbmF0dXJlPTVkYTNmNzI0YzY1NDk0YjQ5NGY4YWZmY2M3Zjc3ZTE4ZTViM2RiY2JmNGI2NDBiOGQ4ZjNlNGIzOGM1MTE0MmY",
			Likes:            userIDs(2),
		},
	}

	if len(components) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}

	return components, nil
}
