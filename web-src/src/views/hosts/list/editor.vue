<template>
    <div class="demo-drawer">
      <tiny-drawer
        v-model:visible="visible"
        :title="fromData.id?'编辑主机':'添加主机'"
        show-footer
        :width="`40%`"
        @close="toggleDrawer(false)"
        @confirm="confirm"
      >
        <tiny-form ref="ruleFormRef" :model="fromData" label-position="left" :label-align="true" :label-width="`150px`" style="margin-top:20px;">
          <tiny-form-item label="主机" prop="host" required>
              <tiny-input v-model="fromData.host"></tiny-input>
          </tiny-form-item>
          <tiny-form-item label="端口号" prop="port" required>
              <tiny-numeric v-model="fromData.port" max="65535" min="1" :controls="false" show-left class="numeric-class"></tiny-numeric>
          </tiny-form-item>
          <tiny-form-item label="用户名" prop="username" required>
              <tiny-input v-model="fromData.username"></tiny-input>
          </tiny-form-item>
          <tiny-form-item label="说明" prop="hostdesc" required>
              <tiny-input v-model="fromData.hostdesc"></tiny-input>
          </tiny-form-item>
          <tiny-form-item label="是否使用SSL证书" prop="iskey" required>
            <tiny-radio-group v-model="fromData.iskey">
              <tiny-radio :label="1">是</tiny-radio>
              <tiny-radio :label="0">否</tiny-radio>
            </tiny-radio-group>
          </tiny-form-item>
          <tiny-form-item label="选择证书" prop="keypath" v-if="fromData.iskey==1">
              <tiny-select v-model="fromData.keypath" clearable>
                <tiny-option label="请选择" value=""> </tiny-option>
                <tiny-option v-for="item in keysList" :key="item.id" :label="item.keyname" :value="item.id"> </tiny-option>
              </tiny-select>
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
    queryKeysList,
  QueryParmas,
} from '@/api/keys';
  import {
    hostsInsert,hostsUpdate
} from '@/api/hosts';

import type {HostsType} from '@/api/hosts';

  const visible = ref<boolean>(false);
  const options = ref<any>([
    {"label":"demo","value":"1"}
  ]);
  const keysList = ref<any[]>([]);
  const fromData = ref<HostsType>({
    id:undefined,
    host:'',
    port:'22',
    username:'root',
    iskey:0,
    keypath:'',
    hostdesc:''
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
  const loadKeysList = ()=>{
    queryKeysList({page:0,limit:0}).then((res:any)=>{
      if(res.code===200){
        keysList.value = res.data;
      }
    });
  }
  const onBeforeClose = (type) => {
    Modal.message({ message: `beforeClose 回调参数 type =  ${type}`, status: 'info', duration: 5000 })
    return false
  }
  
  const confirm = () => {
    let reqData = toRaw(fromData.value)
    reqData.port = reqData.port + "";
      if(reqData.id){
        hostsUpdate(reqData).then((res:any)=>{
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
      }else{
        hostsInsert(reqData).then((res:any)=>{
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
  }
  
   
    // Emits声明
    const emit = defineEmits(["update:visible","update:artdata","success"]);
    watch(propsVisible, (newValue, oldValue) => {
        visible.value = propsVisible.value;
    });
    watch(propsArtdata, (newValue, oldValue) => {
        // @ts-ignore
        fromData.value = propsArtdata.value;
        // @ts-ignore
        delete fromData.value.keyname;
    });
    // 监听值变化，将值传回父组件
    watch(visible, (newValue, oldValue) => {
        emit("update:visible", visible.value);
    });
    watch(fromData, (newValue, oldValue) => {
        emit("update:artdata", fromData.value);
    });
    onMounted(()=>{
      loadKeysList();
    })
</script>
<script lang="ts">
export default {
name: "Editor"
};
</script>