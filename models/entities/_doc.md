## The following schemas are placed inside a nosql db:
Ability -> Ability Sets
Hero Design Tags
[]Mechanism Tags -> Playstyles
Build Paths -> Items


// You may use this alternative method to interpret enum objects
func InterpretEquipmentEnum(i int8) string {
	return [...]string{"Large Wearable", "Small Wearable", "Main Weapon", "Offhand Weapon", "Enchantment", "Bio-Augment"}[i-1]
}