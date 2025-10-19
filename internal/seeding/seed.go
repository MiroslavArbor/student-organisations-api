package seeding

type Seeder interface {
	Seed() error
}

type SeedManager struct {
	seeders []Seeder
}

func NewSeedManager(seeders ...Seeder) *SeedManager {
	return &SeedManager{seeders: seeders}
}

func (m *SeedManager) SeedAll() error {
	for _, seeder := range m.seeders {
		if err := seeder.Seed(); err != nil {
			return err
		}
	}
	return nil
}
