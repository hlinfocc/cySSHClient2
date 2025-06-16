<template>
  <div class="login-form-container">
    <tiny-form
      ref="loginFormInfo"
      :model="loginInfo"
      class="login-form"
      :rules="rules"
      validate-type="text"
      label-width="0"
      size="medium"
    >
      <tiny-form-item prop="username" size="medium">
        <tiny-input
          v-model="loginInfo.username"
          :placeholder="$t('login.form.userName.placeholder')"
        >
        </tiny-input>
      </tiny-form-item>

      <tiny-form-item prop="password" size="medium">
        <tiny-input
          v-model="loginInfo.password"
          type="password"
          show-password
          :placeholder="$t('login.form.password.placeholder')"
        >
        </tiny-input>
      </tiny-form-item>

      <div class="login-form-options">
        <tiny-checkbox v-model="rememberAccountVal" :true-label="1" :false-label="0">{{ $t('login.form.rememberAccount') }}</tiny-checkbox>
        <div>
          <!-- <tiny-link type="primary">
            {{ $t('login.form.forgetPassword') }}
          </tiny-link> -->
          <!-- <tiny-link type="primary" class="divide-line">|</tiny-link> -->
          <!-- <tiny-link type="primary" @click="typeChange">
            {{ $t('login.form.registration') }}
          </tiny-link> -->
        </div>
      </div>

      <tiny-form-item size="medium">
        <tiny-button
          type="info"
          class="login-form-btn"
          :loading="loading"
          @click="handleSubmit"
          >{{ $t('login.form.login') }}</tiny-button
        >
      </tiny-form-item>
    </tiny-form>
  </div>
</template>

<script lang="ts" setup>
  import { inject, ref, reactive, computed, onMounted,onUnmounted } from 'vue';
  import { useRouter } from 'vue-router';
  import {
    Form as TinyForm,
    FormItem as TinyFormItem,
    Input as TinyInput,
    Button as TinyButton,
    Checkbox as TinyCheckbox,
    Link as TinyLink,
    Notify,
    Modal,
  } from '@opentiny/vue';
  import { useI18n } from 'vue-i18n';
  import { useUserStore } from '@/store';
  import useLoading from '@/hooks/loading';
  import { setToken } from '@/utils/auth';

  const router = useRouter();
  const { t } = useI18n();
  const { loading, setLoading } = useLoading();
  const userStore = useUserStore();
  const loginFormInfo = ref();
  const rememberAccountVal = ref(-1);

  const rules = computed(() => {
    return {
      username: [
        {
          required: true,
          message: t('login.form.userName.errMsg'),
          trigger: 'change',
        },
      ],
      password: [
        {
          required: true,
          message: t('login.form.password.errMsg'),
          trigger: 'change',
        },
      ],
    };
  });

  const loginInfo = reactive({
    username: '',
    password: '',
  });

  // 切换模式
  const handle: any = inject('handle');
  const typeChange = () => {
    handle(true);
  };

  function handleSubmit() {
    loginFormInfo.value?.validate(async (valid: boolean) => {
      if (!valid) {
        return;
      }
      if(!import.meta.env.VITE_USE_MOCK){
        window.localStorage.setItem('userRole', 'admin');
        setToken('12345');

        const { redirect, ...othersQuery } = router.currentRoute.value.query;
        router.push({
          name: (redirect as string) || 'Home',
          query: {
            ...othersQuery,
          },
        });
        setLoading(false);
        return
      }
      setLoading(true);

      try {
        
        await userStore.login({
          account: loginInfo.username,
          password: loginInfo.password,
        })
        const { redirect, ...othersQuery } = router.currentRoute.value.query;
        console.log("redirect:>>>>>>",redirect);
        if(rememberAccountVal.value===1){
          window.localStorage.setItem("USER_ACCOUNT_CACHE",loginInfo.username);
        }
        router.push({
          name: redirect?(redirect as string):'Home',
          query: {
            ...othersQuery,
          },
        });
      } catch (err:any) {
        Modal.message({
          message: err.message?err.message:'登录失败',
          status: 'error',
        });
      } finally {
        setLoading(false);
      }
    });
  }
  // 回车执行登录
  const keyDown = (e:any) => {
    if (e.keyCode === 13) {
      handleSubmit();
    }
  };

  onMounted(()=>{
     let accountCache = window.localStorage.getItem("USER_ACCOUNT_CACHE");
     if (accountCache){
      loginInfo.username = accountCache;
     }
     // 回车执行登录
    window.addEventListener('keyup', keyDown, false);
  });
  onUnmounted(() => {
    // 取消监听回车登录
    window.removeEventListener('keyup', keyDown);
  });
</script>

<style lang="less" scoped>
  .login-form-container {
    margin-top: 5%;
  }

  .login-form {
    margin-left: 6%;

    .tiny-form-item {
      margin-bottom: 20px;
    }

    &-container {
      width: 320px;
    }

    &-options {
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 20px;
      font-size: 12px;
    }

    &-btn {
      display: block;
      width: 100%;
      max-width: 100%;
    }
  }

  .divide-line {
    margin: 0 5px;
  }
  // responsive
  @media (max-width: @screen-ms) {
    .login-form {
      margin-left: 5%;

      &-container {
        width: 240px;
      }
    }
  }
</style>
