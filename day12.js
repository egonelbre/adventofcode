function sum(item) {
	if (typeof item === "number") {
		return item;
	}
	if (item instanceof Array) {
		var total = 0;
		for (var i = 0; i < item.length; i++) {
			total += sum(item[i]);
		}
		return total;
	}
	if (typeof item === "object") {
		var total = 0;
		for (var name in item) {
			if (!item.hasOwnProperty(name)) {
				continue;
			}
			if (item[name] === "red") {
				return 0;
			}

			total += sum(item[name]);
		}
		return total;
	}
	return 0
}