<template>
  <v-app >
    <v-navigation-drawer app persistent>
      <LeftNav></LeftNav>
    </v-navigation-drawer>
    <main>
      <v-content>
        <v-container fluid>
          <router-view></router-view>
        </v-container>
      </v-content>
    </main>
  </v-app>
</template>
<script>
import LeftNav from '@/components/LeftNav'
export default {
  name: 'app',
  components: { LeftNav }
}
</script>
