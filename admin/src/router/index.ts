import { createRouter, createWebHashHistory } from "vue-router";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/login",
      name: "login",
      redirect: "/login/index",
      children: [
        {
          path: "index",
          name: "loginIndex",
          component: () => import("@/views/login/Index.vue"),
        },
      ],
    },
    {
      path: "/",
      redirect: "/home",
      component: () => import("@/views/layout/Index.vue"),
      children: [
        {
          path: "home",
          redirect: "/home/index",
          children: [
            {
              path: "index",
              name: "homeIndex",
              component: () => import("@/views/home/Index.vue"),
            },
          ],
        },
        {
          path: "setup",
          redirect: "/setup/basic",
          children: [
            {
              path: "basic",
              name: "setupBasic",
              redirect: "/setup/basic/index",
              children: [
                {
                  path: "index",
                  name: "basicIndex",
                  component: () => import("@/views/basic/Index.vue"),
                },
              ],
            },
            {
              path: "banner",
              redirect: "/setup/banner/index",
              children: [
                {
                  path: "index",
                  name: "bannerIndex",
                  component: () => import("@/views/banner/Index.vue"),
                },
                {
                  path: "info",
                  name: "bannerInfo",
                  component: () => import("@/views/banner/Info.vue"),
                },
              ],
            },
          ],
        },
        {
          path: "column",
          redirect: "/column/index",
          children: [
            {
              path: "index",
              name: "columnIndex",
              component: () => import("@/views/column/Index.vue"),
            },
            {
              path: "info",
              name: "columnInfo",
              component: () => import("@/views/column/Info.vue"),
            },
          ],
        },
      ],
    },
  ],
});

export default router;
