<template>
  <section>
    <m-picker :slots='slots' :datakey='datakey' :valueKey="'v'" :isPicker='isPicker' :indexText='indexText'
    @confirm='pickerConfirm' @cancel='pickerCancel' @change="onValuesChange">
    </m-picker>
  </section>
</template>

<script>
import mPicker from './index';
import areaJson from './china';
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
    },
    isZone: {
      type: Boolean,
      default: false
    },
  },
  data () {
    return {
      indexText: '请选择地区',
        province: areaJson.province,
        city: areaJson.city,
        zone: areaJson.zone,
        oldProvince: '',
        oldCity: '',
      slots: []
    };
  },
  created() {
        let city = this.city[this.province[0].k],
            zone = this.zone[city[0].k];
        if (this.isZone) {
            this.slots = [
                {values: this.province},
                {values: city}
            ];
        } else {
            this.slots = [
                {values: this.province},
                {values: city},
                {values: zone}
            ];
        }
  },
  methods: {
    pickerCancel() {
      this.$emit('cancel');
    },
    pickerConfirm(value, key) {
        let k = '', v = '', zoneK, zoneData = [];
        value.forEach((item)=> {
            k += item.k + '-';
            v += item.v + '-';
        });
        if (this.isZone) { // 当只要省和市的话，confirm第三个参数会 将区作为参数导出，自行操作
            zoneK = value[1].k;
            zoneData = this.zone[zoneK];
            this.$emit('confirm', {k: k.slice(0, -1), v: v.slice(0, -1)}, key, zoneData);
        } else {
            this.$emit('confirm', {k: k.slice(0, -1), v: v.slice(0, -1)}, key);
        }
    },
    onValuesChange(values, picker) {
        let p = values[0].k, c = values[1].k;
        if ((this.oldProvince !== p || this.oldCity !== c)) {
            if (this.oldProvince !== values[0].k) {
               picker.setSlotValues(1, this.city[p]);
               picker.setSlotValues(2, this.zone[this.city[p][0].k]);
            } else if (this.oldCity !== values[0].k) {
                picker.setSlotValues(2, this.zone[c]);
            }
        }
        this.oldProvince = p;
        this.oldCity = c;
    }
  }
};
</script>

<style scoped>
</style>