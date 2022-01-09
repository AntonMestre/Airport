class ValueObject {
    name: String;
    currentValue: Number;
    avgValue: Number;
    unit: String;


    constructor(name: String, currentValue: Number, avgValue: Number, unit: String) {
        this.name = name;
        this.currentValue = currentValue;
        this.avgValue = avgValue;
        this.unit = unit;
    }
}