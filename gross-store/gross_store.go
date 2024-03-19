package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	count, exists := units[unit]
	if !exists {
		return false
	}

	bill[item] = bill[item] + count

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemInBill, itemExists := bill[item]
	if !itemExists {
		return false
	}

	value, valueExists := units[unit]
	if !valueExists {
		return false
	}

	if itemInBill-value < 0 {
		return false
	}

	if itemInBill-value == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = itemInBill - value
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemCountInBill, itemExists := bill[item]
	if !itemExists {
		return 0, false
	}

	return itemCountInBill, true
}
