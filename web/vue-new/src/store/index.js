import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'
import app from './modules/app'
import settings from './modules/settings'
import user from './modules/user'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    app,
    settings,
    user,
    task: {
      namespaced: true,
      state: {
        taskData: {}
      },
      mutations: {
        setTaskData(state, { taskId, data }) {
          state.taskData[taskId] = data
        },
        clearTaskData(state) {
          state.taskData = {}
        }
      },
      actions: {
        setTaskData({ commit }, payload) {
          commit('setTaskData', payload)
        }
      },
      getters: {
        getTaskData: state => taskId => state.taskData[taskId]
      }
    },

    data: {
      namespaced: true,
      state: {
        resData: {}
      },
      mutations: {
        setResData(state, { taskId, data }) {
          state.resData[taskId] = data
        },
        clearResData(state) {
          state.resData = {}
        }
      },
      actions: {
        setResData({ commit }, payload) {
          commit('setResData', payload)
        }
      },
      getters: {
        getResData: state => taskId => state.resData[taskId]
      }
    }
  },
  getters,
  plugins: [createPersistedState({
    // 指定要持久化的模块
    paths: ['task', 'data']
  })]
})

export default store
