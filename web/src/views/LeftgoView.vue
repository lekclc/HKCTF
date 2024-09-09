<template>
  <el-tabs :tab-position="tabPosition" type="card" style="height: 100%" class="demo-tabs" @tab-click="handleClick">
    <el-tab-pane label="Index">
      <div>
        <IndexView />
      </div>
    </el-tab-pane>
    <el-tab-pane label="Game">
      <div>

      </div>
    </el-tab-pane>
    <el-tab-pane label="Post">
      <div>
        
      </div>
    </el-tab-pane>
    <el-tab-pane label="User" >
          
    </el-tab-pane>
    <el-tab-pane label="Config">
      <div>
        
      </div>
    </el-tab-pane>
    
  </el-tabs>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import type { TabsPaneContext } from 'element-plus'
import IndexView from './IndexView.vue';
import UserView from './UserView.vue';
import jwt from '../logic/jwt'

const tabPosition = 'left';
let isAuthenticated = ref(false);
const router = useRouter();

const handleClick = (tab: TabsPaneContext, event: Event) => {
  if (tab.props.label.toLowerCase()==='user') {
      router.push('/login');
      return;
  }
  router.push('/'+tab.props.label.toLowerCase());
}
</script>

<style>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}

.el-tabs--right .el-tabs__content,
.el-tabs--left .el-tabs__content {
  height: 100%;
}
</style>