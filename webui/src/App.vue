<template>
  <div>
      <TopNav></TopNav>
    <main>
      <router-view></router-view>
    </main>
  </div>
</template>
<script>
import TopNav from '@/components/TopNav'
export default {
  name: 'app',
  components: { TopNav }
}
</script>
