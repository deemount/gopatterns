package main

import (
	"fmt"
	"sync"
)

/*

	Beispiel:

	* DBConfig stellt eine Datenbankkonfiguration dar.
	* ConfigManager stellt einen Datenbankkonfigurationsmanager dar.
	* ConfManager ist eine Singleton-Instanz von ConfigManager.
	* GetConfig gibt eine Datenbankkonfiguration für einen bestimmten Schlüssel zurück.
	* SetConfig setzt eine Datenbankkonfiguration für einen bestimmten Schlüssel.

	In diesem Beispiel habe ich einen Datenbankkonfigurationsmanager
	dargestellt durch die ConfigManager-Struktur.

	Der Manager enthält eine Karte von Datenbankkonfigurationen,
	wobei jede Konfiguration durch einen eindeutigen Schlüssel identifiziert wird.
	Der sync.Once-Typ wird verwendet, um sicherzustellen,
	dass die Abbildung der Konfigurationen nur einmal erstellt wird, auch wenn
	mehrere Goroutinen gleichzeitig versuchen, auf die Map zuzugreifen.

	Die Instanzvariable ist eine Singleton-Instanz des ConfigManager.
	Die Methoden GetConfig und SetConfig ermöglichen das Abrufen und
	das Setzen von Datenbankkonfigurationen in der Map.

*/

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

type ConfigManager struct {
	configs map[string]*DBConfig
	once    sync.Once
}

var ConfManager ConfigManager

func (c *ConfigManager) GetConfig(key string) (*DBConfig, error) {
	c.once.Do(func() {
		c.configs = make(map[string]*DBConfig)
	})
	config, ok := c.configs[key]
	if !ok {
		return nil, fmt.Errorf("config for key '%s' not found", key)
	}
	return config, nil
}

func (c *ConfigManager) SetConfig(key string, config *DBConfig) {
	c.once.Do(func() {
		c.configs = make(map[string]*DBConfig)
	})
	c.configs[key] = config
}
