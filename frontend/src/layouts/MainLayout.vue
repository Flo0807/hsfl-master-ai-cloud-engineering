<template>
  <q-layout view="hHh Lpr lff">
    <q-header elevated class="bg-primary text-white">
      <q-toolbar>
        <q-toolbar-title>
          <span>BoardHub</span>
        </q-toolbar-title>
        <q-btn color="red" label="Logout" @click="logout"/>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import {useRouter} from 'vue-router';
import {useAuthStore} from 'stores/auth-store';

const router = useRouter()
const authStore = useAuthStore()

function logout() {
  authStore.logout();
  router.push({name: 'login'})
}
</script>
