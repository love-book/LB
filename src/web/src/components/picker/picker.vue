<template>
  <div class="picker" v-show='isPicker' :class="{ 'picker-3d': rotateEffect }">
    <div class="picker-toolbar" v-if="showToolbar">
      <button class="left" @click="cancel">取消</button>
      <span>{{indexText}}</span>
      <button class="right" @click="confirm">确定</button>
    </div>
    <div class="picker-items">
      <picker-slot v-for="slot in slots" :valueKey='valueKey' :values="slot.values || []" :text-align="slot.textAlign || 'center'" :visible-item-count="visibleItemCount" :class-name="slot.className" :flex="slot.flex || '1'" v-model="values[slot.valueIndex]" :rotate-effect="rotateEffect" :divider="slot.divider" :content="slot.content"></picker-slot>
      <div class="picker-center-highlight"></div>
    </div>
  </div>
</template>

<script>
  import PickerSlot from './picker-slot.vue';
  export default {
    name: 'mt-picker',

    componentName: 'picker',

    props: {
      slots: {
        type: Array
      },
      showToolbar: {
        type: Boolean,
        default: true
      },
      indexText: {
        type: String,
        default: '请选择'
      },
      isPicker: {
        type: Boolean,
        default: false
      },
      visibleItemCount: {
        type: Number,
        default: 5
      },
      rotateEffect: {
        type: Boolean,
        default: false
      },
      valueKey: {
        type: String,
        default: ''
      },
      datakey: {
        type: String,
        default: ''
      }
    },

    created() {
      this.$on('slotValueChange', this.slotValueChange);
      var slots = this.slots || [];
      this.values = [];
      var values = this.values;
      var valueIndexCount = 0;
      slots.forEach(function(slot) {
        if (!slot.divider) {
          slot.valueIndex = valueIndexCount++;
          values[slot.valueIndex] = (slot.values || [])[slot.defaultIndex || 0];
        }
      });
    },

    methods: {
      slotValueChange() {
        this.$emit('change', this.values, this);
      },
      cancel() {
        this.$emit('cancel');
      },
      confirm() {
        if (this.values.undefined) {
          this.$emit('confirm', this.values.undefined, this.datakey, this);
        } else {
          this.$emit('confirm', this.values, this.datakey, this);
        }
      },

      getSlot(slotIndex) {
        var slots = this.slots || [];
        var count = 0;
        var target;
        var children = this.$children;

        slots.forEach(function(slot, index) {
          if (!slot.divider) {
            if (slotIndex === count) {
              target = children[index];
            }
            count++;
          }
        });

        return target;
      },
      getSlotValue(index) {
        var slot = this.getSlot(index);
        if (slot) {
          return slot.value;
        }
        return null;
      },
      setSlotValue(index, value) {
        var slot = this.getSlot(index);
        if (slot) {
          slot.currentValue = value;
        }
      },
      getSlotValues(index) {
        var slot = this.getSlot(index);
        if (slot) {
          return slot.mutatingValues;
        }
        return null;
      },
      setSlotValues(index, values) {
        var slot = this.getSlot(index);
        if (slot) {
          slot.mutatingValues = values;
        }
      },
      getValues() {
        return this.values;
      },
      setValues(values) {
        var slotCount = this.slotCount;
        values = values || [];
        if (slotCount !== values.length) {
          throw new Error('values length is not equal slot count.');
        }
        values.forEach((value, index) => {
          this.setSlotValue(index, value);
        });
      }
    },

    computed: {
      // values() {
      //   var slots = this.slots || [];
      //   var values = [];
      //   slots.forEach(function(slot) {
      //     if (!slot.divider) values.push(slot.values);
      //   });
      //   return values;
      // },
      slotCount() {
        var slots = this.slots || [];
        var result = 0;
        slots.forEach(function(slot) {
          if (!slot.divider) result++;
        });
        return result;
      }
    },

    components: {
      PickerSlot
    }
  };
</script>
<style>
  .picker {
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    overflow: hidden;
    background: #fff;
    border-top: 1px solid #f0f0f0;
    animation:fadeInUp .4s .2s ease both;
    -webkit-animation:fadeInUp .4s .2s ease both;
    -moz-animation:fadeInUp .4s .2s ease both;
    -o-animation:fadeInUp .4s .2s ease both;
  }

  .picker-toolbar {
    height: 40px;
  }

  .picker-items {
    display: -webkit-box;
    display: -ms-flexbox;
    display: -webkit-flex;
    display: flex;
    -webkit-box-pack: center;
    -ms-flex-pack: center;
    -webkit-justify-content: center;
    justify-content: center;
    padding: 0;
    text-align: right;
    font-size: 24px;
    position: relative;
  }

  .picker-center-highlight {
    height: 36px;
    box-sizing: border-box;
    position: absolute;
    left: 0;
    width: 100%;
    top: 50%;
    margin-top: -18px;
    pointer-events: none
  }

  .picker-center-highlight:before,
  .picker-center-highlight:after {
    content: '';
    position: absolute;
    height: 1px;
    width: 100%;
    background-color: #eaeaea;
    display: block;
    z-index: 15;
    -webkit-transform: scaleY(.5);
    transform: scaleY(0.5);
  }

  .picker-center-highlight:before {
    left: 0;
    top: 0;
    bottom: auto;
    right: auto;
  }

  .picker-center-highlight:after {
    left: 0;
    bottom: 0;
    right: auto;
    top: auto;
  }
  @-webkit-keyframes fadeInUp{
  0%{
    opacity:0;
    transform:translateY(50px);
    -webkit-transform:translateY(50px);
    -moz-transform:translateY(50px);
    -o-transform:translateY(50px);
  }
  100%{
    opacity:1;
    transform:translateY(0);
    -webkit-transform:translateY(0);
    -moz-transform:translateY(0);
    -o-transform:translateY(0)
    }
  }
@-moz-keyframes fadeInUp{
  0%{
    opacity:0;
    transform:translateY(50px);
    -webkit-transform:translateY(50px);
    -moz-transform:translateY(50px);
    -o-transform:translateY(50px);
  }
100%{
    opacity:1;
    transform:translateY(0);
    -webkit-transform:translateY(0);
    -moz-transform:translateY(0);
    -o-transform:translateY(0)
  }
}
  button{
  -webkit-appearance: none;
    -moz-appearance: none;
    appearance: none;
    border-radius: 4px;
    border: 0;
    box-sizing: border-box;
    color: inherit;
    display: block;
    height: .8rem;
  font-size: .28rem;
    outline: 0;
    overflow: hidden;
    position: relative;
    text-align: center;
  color: #fff;
}
.picker-toolbar{
  position: relative;
}
.picker-toolbar button{
  padding-right: .4rem;
  background: none;
  color: #333;
  display: inline-block;
}
.picker-toolbar button.left{
  position: absolute;
  top: 0;
  left: 0;
  padding-left: .4rem;
  padding-left: .4rem;
}
.picker-toolbar button.right{
  position: absolute;
  top: 0;
  right: 0;
  padding-left: .4rem;
  padding-right: .4rem;
}
.picker-toolbar span{
  display: block;
  width: 70%;
  text-align: center;
  margin: 0 auto;
  height: .8rem;
  line-height: .8rem;
}
</style>

