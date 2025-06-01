<template>
    <div class="demo-drawer">
      <tiny-drawer
        v-model:visible="visible"
        :title="$t('keys.form.title')"
        show-footer
        :width="`40%`"
        @close="toggleDrawer(false)"
        @confirm="confirm"
      >
        <tiny-form ref="ruleFormRef" :model="fromData" label-position="left" :label-align="true" :label-width="`150px`" style="margin-top:20px;">
          <tiny-form-item :label="$t('keys.form.name')" prop="keyname" required>
              <tiny-input v-model="fromData.keyname" :placeholder="$t('keys.form.name.placeholder')"></tiny-input>
          </tiny-form-item>
          <tiny-form-item :label="$t('keys.form.passwd')" prop="passwd">
              <tiny-input v-model="fromData.passwd" type="password" :placeholder="$t('keys.form.passwd.placeholder')"></tiny-input>
          </tiny-form-item>
          
        </tiny-form>
      </tiny-drawer>
    </div>
  </template>
  <script lang="ts" setup>
  import { ref, defineProps, toRef, watch, onMounted, toRaw } from "vue";
  import { 
    Drawer as TinyDrawer, 
    Button as TinyButton, 
    Modal,
    Form as TinyForm,
    FormItem as TinyFormItem,
    Input as TinyInput,
    Select as TinySelect, 
    Option as TinyOption,
    TinyRadio, TinyRadioGroup,TinyNumeric,TinyNotify
   } from '@opentiny/vue';
   
  import {
    keysGenerate,
} from '@/api/keys';


import type {generateParmas} from '@/api/keys';

  const visible = ref<boolean>(false);
  const options = ref<any>([
    {"label":"demo","value":"1"}
  ]);
  const keysList = ref<any[]>([]);
  const fromData = ref<generateParmas>({
    keyname:'',
    passwd:'',
  });
  const props = defineProps({
    visible: {
        type: Boolean,
        default: false
    },
    artdata: {
        type: Object,
        default: ()=>{}
    }
   });
   const propsVisible = toRef(props, "visible");
   const propsArtdata = toRef(props, "artdata");
  const toggleDrawer = (value) => {
    visible.value = value
  }

  const onBeforeClose = (type) => {
    Modal.message({ message: `beforeClose 回调参数 type =  ${type}`, status: 'info', duration: 5000 })
    return false
  }
  
  const confirm = () => {
    console.log("fromData:",fromData.value);
    keysGenerate(fromData.value).then((res:any)=>{
      let noteType = 'error';
      if(res.code===200){
        visible.value = false;
        noteType = 'success';
        emit("success", true);
      }
      TinyNotify({
        type: noteType,
        message: res.msg,
        position: 'top-right'
      })
    })
  }
  
   
    // Emits声明
    const emit = defineEmits(["update:visible","update:artdata","success"]);
    watch(propsVisible, (newValue, oldValue) => {
        visible.value = propsVisible.value;
    });
   
    // 监听值变化，将值传回父组件
    watch(visible, (newValue, oldValue) => {
        emit("update:visible", visible.value);
    });

    onMounted(()=>{
    })
</script>
<script lang="ts">
export default {
name: "Editor"
};
</script>