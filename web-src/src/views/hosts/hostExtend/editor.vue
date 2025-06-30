<template>
    <div class="demo-drawer">
      <tiny-drawer
        v-model:visible="visible"
        :title="$t('hostExtend.form.title')"
        show-footer
        :width="`40%`"
        @close="toggleDrawer(false)"
        @confirm="confirm"
      >
        <tiny-form ref="ruleFormRef" :model="fromData" label-position="left" :label-align="true" :label-width="`190px`" style="margin-top:20px;">
          <tiny-form-item :label="$t('hostExtend.form.ID')" prop="id">
              <tiny-select v-model="fromData.id" clearable>
                <tiny-option label="请选择" value=""> </tiny-option>
                <tiny-option v-for="item in hostList" :key="item.id" :label="item.host" :value="item.id"> </tiny-option>
              </tiny-select>
          </tiny-form-item>
          <tiny-form-item :label="$t('hostExtend.columns.cloudType')" prop="cloudType" required>
              <tiny-input v-model="fromData.cloudType" :placeholder="$t('hostExtend.form.cloudType.placeholder')"></tiny-input>
          </tiny-form-item>
          <tiny-form-item :label="$t('hostExtend.columns.startTime')" prop="startTime" required>
            <tiny-date-picker v-model="fromData.startTime" type="datetime" value-format="yyyy-MM-dd HH:mm:ss" :isutc8="true"></tiny-date-picker>
          </tiny-form-item>
          <tiny-form-item :label="$t('hostExtend.columns.endTime')" prop="endTime" required>
            <tiny-date-picker v-model="fromData.endTime" type="datetime" value-format="yyyy-MM-dd HH:mm:ss" :isutc8="true"></tiny-date-picker>
          </tiny-form-item>
          <tiny-form-item :label="$t('hostExtend.columns.izCrond')" prop="izCrond" required :extra="$t('hostExtend.form.izCrond.extra')">
            <tiny-radio-group v-model="fromData.izCrond">
              <tiny-radio :label="1">{{ $t('hostExtend.columns.izCrond.yes') }}</tiny-radio>
              <tiny-radio :label="0">{{ $t('hostExtend.columns.izCrond.no') }}</tiny-radio>
            </tiny-radio-group>
          </tiny-form-item>
          <tiny-form-item :label="$t('hostExtend.columns.remarks')" prop="remarks">
              <tiny-input v-model="fromData.remarks" :placeholder="$t('hostExtend.form.remarks.placeholder')"></tiny-input>
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
    TinyRadio, TinyRadioGroup,TinyNumeric,TinyNotify,TinyDatePicker
   } from '@opentiny/vue';
   import {
  queryHostsList,
  QueryParmas,
} from '@/api/hosts';
import type { HostsExtendType } from '@/api/hostsExtend';
import { hostExtendInsert, hostExtendUpdate } from '@/api/hostsExtend';

  const visible = ref<boolean>(false);
  const izUpdate = ref<boolean>(false);
  const options = ref<any>([
    {"label":"demo","value":"1"}
  ]);
  const hostList = ref<any[]>([]);
  const fromData = ref<HostsExtendType>({
    cloudType:'',
    startTime:'',
    endTime:'',
    izCrond:0,
    remarks:''
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
    if(!izUpdate.value){
      hostExtendInsert(fromData.value).then((res:any)=>{
        let noteType = 'error';
        if(res.code===200){
          visible.value = false;
          noteType = 'success';
          emit("success", true);
          emit("update:artdata", fromData.value);
          emit("update:visible", false);
        }
        TinyNotify({
          type: noteType,
          message: res.msg,
          position: 'top-right'
        })
      })
    }else{
      hostExtendUpdate(fromData.value).then((res:any)=>{
        let noteType = 'error';
        if(res.code===200){
          visible.value = false;
          noteType = 'success';
          emit("success", true);
          emit("update:artdata", fromData.value);
          emit("update:visible", false);
        }
        TinyNotify({
          type: noteType,
          message: res.msg,
          position: 'top-right'
        })
      })
    }
  }
  
  const loadHosts = ()=>{
    const queryParmas = {
      page: 0,
      limit: 0,
      hostExtent: 1,
  };
    queryHostsList(queryParmas).then((res:any)=>{
        if(res.code === 200){
          hostList.value = res.data;
        }
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
    watch(propsArtdata, (newValue, oldValue) => {
        // @ts-ignore
        fromData.value = propsArtdata.value;
        // @ts-ignore
        izUpdate.value = true;
    });

    onMounted(()=>{
      loadHosts();
    })
</script>
<script lang="ts">
export default {
name: "Editor"
};
</script>