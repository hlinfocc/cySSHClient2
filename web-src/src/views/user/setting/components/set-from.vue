<template>
  <tiny-layout>
    <tiny-form
      ref="setFormRef"
      :model="state.filterOptions"
      :rules="rules"
      label-width="150px"
      :label-align="true"
      label-position="left"
      size="small"
    >
      <tiny-row :flex="true" justify="left">
        <tiny-col :span="5" label-width="100px">
          <tiny-form-item
            :label="$t('userSetting.realName')" prop="realName"
          >
            <tiny-input v-model="state.filterOptions.realName"></tiny-input>
          </tiny-form-item>
        </tiny-col>
        <tiny-col :span="5" label-width="100px">
          <tiny-form-item :label="$t('userSetting.account')" prop="account">
            <tiny-input v-model="state.filterOptions.account" disabled></tiny-input>
          </tiny-form-item>
        </tiny-col>
      </tiny-row>

      <!-- <tiny-row :flex="true" justify="left">
        <tiny-col :span="5" label-width="100px">
          <tiny-form-item :label="$t('userSetting.status')" prop="status">
            <tiny-input
              v-model="state.filterOptions.status"
            ></tiny-input>
          </tiny-form-item>
        </tiny-col>
        <tiny-col :span="5" label-width="100px">
          <tiny-form-item :label="$t('userSetting.userType')" prop="userType">
            <tiny-input
              v-model="state.filterOptions.userType"
            ></tiny-input>
          </tiny-form-item>
        </tiny-col>
      </tiny-row> -->

      <tiny-row :flex="true" justify="left">
        <tiny-col :span="5" label-width="100px">
          <tiny-form-item :label="$t('userSetting.lastLoginTime')">
            <tiny-input
              v-model="state.filterOptions.lastLoginTime" disabled
            ></tiny-input>
          </tiny-form-item>
        </tiny-col>
        <tiny-col :span="5" label-width="100px">
          <tiny-form-item :label="$t('userSetting.lastLoginIp')">
            <tiny-input
              v-model="state.filterOptions.lastLoginIp" disabled
            ></tiny-input>
          </tiny-form-item>
        </tiny-col>
      </tiny-row>

      <tiny-row :flex="true" justify="left">
        <tiny-col :span="5" label-width="100px">
          <tiny-form-item :label="$t('userSetting.role')">
            <tiny-input v-model="state.filterOptions.role" disabled></tiny-input>
          </tiny-form-item>
        </tiny-col>
        <tiny-col :span="5" label-width="100px">
          <tiny-form-item :label="$t('userSetting.passwd')">
            <tiny-input
              v-model="state.filterOptions.passwd"
              type="password"
              show-password
            ></tiny-input>
          </tiny-form-item>
        </tiny-col>
      </tiny-row>
    </tiny-form>
  </tiny-layout>
</template>

<script lang="ts" setup>
  import { ref, reactive, defineProps, computed, defineExpose, onMounted, toRaw } from 'vue';
  import { useI18n } from 'vue-i18n';
  import {
    Select as TinySelect,
    Option as TinyOption,
    Layout as TinyLayout,
    Form as TinyForm,
    FormItem as TinyFormItem,
    Row as TinyRow,
    Col as TinyCol,
    Input as TinyInput,
    DatePicker as TinyDatePicker,
    Modal,
  } from '@opentiny/vue';
  import { useUserStore } from '@/store';

  const userStore = useUserStore();

  interface FilterOptions {
    realName: string;
    account: string;
    status: number;
    userType: number;
    role: string;
    lastLoginTime: string;
    lastLoginIp: string;
    createTime: string;
    updateTime: string;
    passwd: string;
  }

  const projectData = [
    {
      value: '1',
      label: 'social recruitment',
    },
    {
      value: '2',
      label: 'scholl recruitment',
    },
    {
      value: '3',
      label: 'Job transfer',
    },
  ];

  // 加载效果
  const state = reactive<{
    filterOptions: FilterOptions;
    department: string;
    position: Array<object>;
    type: Array<object>;
    date: Array<object>;
    during: string;
    startTime: string;
    endTime: string;
  }>({
    filterOptions: {} as FilterOptions,
    department: '',
    position: [],
    type: [],
    date: [],
    during: '',
    startTime: '',
    endTime: '',
  });

  // 初始化请求数据
  const setFormRef = ref();
  const { t } = useI18n();

  // 校验规则
  const rulesType = {
    required: true,
    trigger: 'blur',
  };
  const rulesSelect = {
    required: true,
    message: '必选',
    trigger: 'blur',
  };
  const rules = computed(() => {
    return {
      department: [rulesType],
      position: [rulesType],
      type: [rulesSelect],
      date: [rulesType],
      during: [rulesType],
      startTime: [rulesType],
      endTime: [rulesType],
    };
  });

  
  const setFormValid = async() => {
    let setValidate = false;
    setFormRef.value.validate(async(valid: boolean) => {
      setValidate = valid;
      if (valid){
        let formData= toRaw(state.filterOptions);
        console.log("formdata:",formData);
        await userStore.updateInfo(formData).then((res:any)=>{
          console.log(">>>>>_____<<>>>>:",res);
          Modal.message({
            message: res.msg,
            status: res.code===200?'success':'error',
          });
        });
      }
    });

    return setValidate;
  };

  const setReset = () => {
    state.filterOptions = {} as FilterOptions;
  };

  const setData = () => {
    return state;
  };

  onMounted(()=>{
    let user:any = window.localStorage.getItem('USER_INFO_STORE_STATE');
    state.filterOptions = JSON.parse(user);
  })

  defineExpose({
    setData,
    setFormValid,
    setReset,
  });
</script>

<style scoped lang="less">
  :deep(.tiny-row) {
    margin-bottom: 15px;
  }
</style>
