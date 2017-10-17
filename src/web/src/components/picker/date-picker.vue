<template>
	<section>
		<m-picker :slots='slots' :datakey='datakey' :isPicker='isPicker' :indexText='indexText'
		@confirm='pickerConfirm' @cancel='pickerCancel' @change="onValuesChange">
		</m-picker>
	</section>
</template>

<script>
import mPicker from './index';
export default {
  name: 'address-picker',
  components: {
  	mPicker
  },
  props: {
  	isPicker: {
  		type: Boolean,
  		default: false
  	},
  	datakey: {
  		type: String,
  		default: ''
  	}
  },
  data () {
    return {
    	indexText: '请选择日期',
    	oldValue: '',
    	slots: []
    };
  },
  created() {
	  	this.slots = [
	  		{values: this.createArray(2017, 2030)},
	  		{values: this.createArray(1, 12)},
	  		{values: this.createArray(1, 31)}
	  	];
  },
  methods: {
  	createArray(min, max) {
  		let newArray = [];
  		for (let i = min; i <= max; i++) {
  			i = i < 10 ? `0${i}` : i;
  			newArray.push(i);
  		}
  		return newArray;
  	},
  	pickerCancel() {
  		this.$emit('cancel');
  	},
  	pickerConfirm(value, key) {
  		this.$emit('confirm', value, key);
  	},
  	getDaysInMonth(year, month){
        month = parseInt(month, 10);
        let temp = new Date(year, month, 0), newArray = [];
        for(let i = 1; i <= temp.getDate(); i++) {
        	i = i < 10 ? `0${i}` : i;
        	newArray.push(i);
        }
        return newArray;
    },
  	onValuesChange(values, picker) {
  		if (this.oldValue && this.oldValue !== values[0]+values[1]) {
  			picker.setSlotValues(2, this.getDaysInMonth(values[0], values[1]));
  		};
  		this.oldValue = values[0]+values[1];
    }
  }
};
</script>

<style scoped>
</style>