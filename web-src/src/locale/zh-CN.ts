import localeLogin from '@/views/login/locale/zh-CN';
import localeTheme from '@/components/theme/locale/zh-CN';

import localeSearchTable from '@/views/list/search-table/locale/zh-CN';

import localeHostsList from '@/views/hosts/list/locale/zh-CN';
import localeKeysList from '@/views/keys/list/locale/zh-CN';
import localeHostExtendList from '@/views/hosts/hostExtend/locale/zh-CN';

import locale403 from '@/views/exception/403/locale/zh-CN';
import locale404 from '@/views/exception/404/locale/zh-CN';
import locale500 from '@/views/exception/500/locale/zh-CN';

import localeUserInfo from '@/views/user/info/locale/zh-CN';
import localeUserSetting from '@/views/user/setting/locale/zh-CN';

import localekanban from '@/views/board/locale/zh-CN';


import localeSettings from './zh-CN/settings';

import localeHttpError from './zh-CN/httpError';

export default {
  'menu.board': '首页',
  'menu.home': '欢迎',
  'menu.work': '工作台',
  'menu.list': '列表页',
  'menu.hosts': '主机管理',
  'menu.result': '结果页',
  'menu.exception': '异常页',
  'menu.notFound': '未知页面',
  'menu.keys': '密钥管理',
  'menu.profile.detail': '基础详情页',
  'menu.visualization': '数据可视化',
  'menu.user': '个人中心',
  'navbar.docs': '文档中心',
  'navbar.action.locale': '切换为中文',
  'messageBox.switchRoles': '切换角色',
  'messageBox.userCenter': '用户中心',
  'messageBox.userSettings': '用户设置',
  'messageBox.logout': '退出登录',
  'menu.cloud': '云服务能力展示',
  ...localeTheme,
  ...localeSettings,
  ...localeLogin,
  ...localeSearchTable,
  ...locale403,
  ...locale404,
  ...locale500,
  ...localeUserInfo,
  ...localeUserSetting,
  ...localekanban,
  ...localeHttpError,
  ...localeHostsList,
  ...localeKeysList,
  ...localeHostExtendList,
};
