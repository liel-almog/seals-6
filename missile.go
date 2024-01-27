package main

type launcher string

const (
	torpedoLauncher    launcher = "Torpedo"
	ballisticLauncher  launcher = "Balistic"
	cruiseLauncher     launcher = "Cruise"
	hypersonicLauncher launcher = "Hypersonic"
)

type missileLauncher interface {
	launch()
	addMissiles(count int)
	clearMissiles()
	len() int
}

type missile struct {
	failed bool
}

type missileStorage struct {
	missiles []missile
}

func (m *missileStorage) addMissiles(count int) {
	for i := 0; i < count; i++ {
		m.missiles = append(m.missiles, missile{})
	}
}

func (m *missileStorage) clearMissiles() {
	m.missiles = []missile{}
}

func (m *missileStorage) len() int {
	return len(m.missiles)
}

func newMissileStorage() *missileStorage {
	return &missileStorage{
		missiles: []missile{},
	}
}

type torpedoMissileLauncher struct {
	*missileStorage
}

func (t *torpedoMissileLauncher) launch() {
	println("Torpedo launched!")
}

type ballisticMissileLauncher struct {
	*missileStorage
}

func (b *ballisticMissileLauncher) launch() {
	println("Ballistic missile launched!")
}

type cruiseMissileLauncher struct {
	*missileStorage
}

func (c *cruiseMissileLauncher) launch() {
	println("Cruise missile launched!")
}

type hypersonicMissileLauncher struct {
	*missileStorage
}

func (h *hypersonicMissileLauncher) launch() {
	println("Hypersonic missile launched!")
}

var launchers = map[launcher]missileLauncher{
	torpedoLauncher: &torpedoMissileLauncher{
		missileStorage: newMissileStorage(),
	},
	ballisticLauncher: &ballisticMissileLauncher{
		missileStorage: newMissileStorage(),
	},
	cruiseLauncher: &cruiseMissileLauncher{
		missileStorage: newMissileStorage(),
	},
	hypersonicLauncher: &hypersonicMissileLauncher{
		missileStorage: newMissileStorage(),
	},
}

func initLaunchers() {
	for _, l := range launchers {
		l.addMissiles(10)
	}
}
