package discovery

import "pg-verify/internal/models"

type MockDiscovery struct{}

func NewMockDiscovery() *MockDiscovery {
	return &MockDiscovery{}
}

func (m *MockDiscovery) Discover() ([]models.Cluster, error) {
	clusters := []models.Cluster{
		{
			Name:      "payments-db",
			Namespace: "finance",
		},
		{
			Name:      "users-db",
			Namespace: "default",
		},
		{
			Name:      "analytics-db",
			Namespace: "analytics",
		},
	}

	return clusters, nil
}
