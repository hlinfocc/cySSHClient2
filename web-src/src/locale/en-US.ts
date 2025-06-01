import localeLogin from '@/views/login/locale/en-US';
import localeTheme from '@/components/theme/locale/en-US';


import locale403 from '@/views/exception/403/locale/en-US';
import locale404 from '@/views/exception/404/locale/en-US';
import locale500 from '@/views/exception/500/locale/en-US';

import localeUserInfo from '@/views/user/info/locale/en-US';
import localeUserSetting from '@/views/user/setting/locale/en-US';

import localekanban from '@/views/board/locale/en-US';

import localeHostsList from '@/views/hosts/list/locale/en-US';
import localeKeysList from '@/views/keys/list/locale/en-US';
import localeHostExtendList from '@/views/hosts/hostExtend/locale/en-US';

import localeSettings from './en-US/settings';

import localeHttpError from './en-US/httpError';



export default {
  'menu.board': 'Dashboard Page',
  'menu.home': 'Monitoring page',
  'menu.work': 'workbench',
  'menu.list': 'List',
  'menu.hosts': 'Host Manager',
  'menu.result': 'Result',
  'menu.exception': 'Exception',
  'menu.form': 'Form',
  'menu.keys': 'keys for ssh',
  'menu.profile.detail': 'Basic details page',
  'menu.visualization': 'Data Visualization',
  'menu.user': 'User Center',
  'navbar.docs': 'Docs',
  'navbar.action.locale': 'Switch to English',
  'messageBox.switchRoles': 'Switch Roles',
  'messageBox.userCenter': 'User Center',
  'messageBox.userSettings': 'User Settings',
  'messageBox.logout': 'Logout',
  'menu.cloud': 'Cloud service capability',
  ...localeTheme,
  ...localeSettings,
  ...localeLogin,
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
