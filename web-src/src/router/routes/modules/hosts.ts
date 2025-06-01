import { RoleType } from '@/types/roleType';

export default {
  path: 'hosts',
  name: 'Hosts',
  id: 'Hosts',
  label: 'Hosts',
  component: () => import('@/views/hosts/index.vue'),
  meta: {
    locale: 'menu.hosts',
    requiresAuth: true,
    order: 2,
    roles: [RoleType.admin, RoleType.user],
  },
  children: [
    {
      path: 'hostsList', // The midline path complies with SEO specifications
      name: 'HostsList',
      id: 'HostsList',
      label: 'HostsList',
      component: () => import('@/views/hosts/list/index.vue'),
      meta: {
        locale: 'menu.hosts.list',
        requiresAuth: true,
        roles: [RoleType.admin, RoleType.user],
      },
    },
    {
      path: 'hostsExtend', // The midline path complies with SEO specifications
      name: 'HostsExtend',
      id: 'HostsExtend',
      label: 'HostsExtend',
      component: () => import('@/views/hosts/hostExtend/index.vue'),
      meta: {
        locale: 'menu.hostExtend.list',
        requiresAuth: true,
        roles: [RoleType.admin, RoleType.user],
      },
    },
  ],
};
