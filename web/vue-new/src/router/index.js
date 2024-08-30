import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'
import axios from 'axios';

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'/'el-icon-x' the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/',
    component: () => import('@/views/login/index'),
    // component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/register',
    component: () => import('@/views/register/index'),
    hidden: true
  },

  {
    path: '/users',
    component: () => import('@/views/user/userList.vue'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/dashboard',
    component: Layout,
    redirect: '/dashboard/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '首页', icon: 'dashboard' }
    }]
  },
 
  // {
  //   path: '/log',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'index',
  //       name: 'log',
  //       component: () => import('@/views/log/index'),
  //       meta: { title: '日志管理', icon: 'form' }
  //     }
  //   ]
  // },
  {
    path: '/sys',
    component: Layout,
    redirect: '/sys/user',
    name: 'sysManage',
    meta: { title: '系统管理', icon: 'sys' },
    children: [
      {
        path: 'user',
        name: 'user',
        component: () => import('@/views/sys/user'),
        meta: { title: '个人中心', icon: 'userManage' }
      },
      {
        path: 'role',
        name: 'role',
        component: () => import('@/views/sys/role'),
        meta: { title: '密码修改', icon: 'roleManage' }
      },
      // {
      //   path: 'config',
      //   name: 'config',
      //   component: () => import('@/views/sys/config'),
      //   meta: { title: '相关配置', icon: 'sys' }
      // }
    ]
  },
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

router.beforeEach((to, from, next) => {
  const data = { role: 0 }; // 定义一个包含 role 属性的对象

  const token = window.sessionStorage.getItem('token');
  if (token) {
    axios.get('parse/token', {
        headers: {
          'Authorization': `Bearer ${token}`, // 将JWT令牌添加到请求头中
          'Content-Type': 'application/json' // 设置请求头为JSON类型
        }
      })
      .then(res => {
        // 从响应数据中获取角色，并存储到对象中
        data.role = res.data.role;
        // 继续执行路由跳转
        continueRouting();
      })
      .catch(error => {
        console.error('Error fetching data:', error); // 处理请求错误
        // 继续执行路由跳转
        continueRouting();
      });
  } else {
    // 如果没有token，则直接执行路由跳转
    continueRouting();
  }

  function continueRouting() {
    // 根据用户角色进行路由控制
    if (token) {
      if (data.role !== '2' && to.path === '/users') {
        alert('用户无权限');
        next(from.path);
      } else {
        next();
      }
    } else if (to.path === '/register' || to.path === '/') {
      next();
    } else if (to.path !== '/') {    // 如果用户未登录且访问其他页面，则重定向到登录页面
      next('/');
    } else {
      next();
    }
  }
});

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router

