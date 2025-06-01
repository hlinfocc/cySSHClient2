import { RoleType } from '@/types/roleType';

export default {
  path: 'keys',
  name: 'Keys',
  id: 'Keys',
  label: 'Keys',
  component: () => import('@/views/keys/list/index.vue'),
  meta: {
    locale: 'menu.keys',
    requiresAuth: true,
    order: 3,
    roles: [RoleType.admin, RoleType.user],
  }
};
