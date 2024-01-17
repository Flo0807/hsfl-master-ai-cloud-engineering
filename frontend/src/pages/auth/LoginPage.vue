<script lang="ts" setup>
import {ref} from 'vue';
import {useRoute, useRouter} from 'vue-router';
import {useAuthStore} from 'stores/auth-store';
import {useQuasar} from 'quasar';

const $q = useQuasar()
const route = useRoute();
const router = useRouter()
const authStore = useAuthStore();

const email = ref('')
const password = ref('')

const goToRegister = () => {
  router.push({name: 'register'})
}

const resetInputs = () => {
  email.value = '';
  password.value = '';
}

const login = () => {
  authStore.login(
    {
      email: email.value,
      password: password.value
    }
  ).then(
    () => {
      resetInputs();
      const nextPath = Array.isArray(route.query.next) ? route.query.next[0] : route.query.next;
      router.push({ path: nextPath ?? '/' });
    }
  ).catch(() => {
    $q.notify({
      color: 'negative',
      position: 'top',
      message: 'Login failed',
      icon: 'report_problem'
    })
  })
}
</script>

<template>
  <q-page class="flex flex-center bg-grey-2">
    <q-card bordered class="q-pa-md shadow-2 my_card">
      <q-card-section class="text-center">
        <div class="text-grey-9 text-h5 text-weight-bold">Sign in</div>
        <div class="text-grey-8">Sign in below to access your account</div>
      </q-card-section>

      <!-- Form start -->
      <q-form @submit.prevent="login">
        <q-card-section>
          <q-input v-model="email" dense label="Email Address" outlined type="email"></q-input>
          <q-input v-model="password" class="q-mt-md" dense label="Password" outlined type="password"></q-input>
        </q-card-section>

        <q-card-section>
          <q-btn class="full-width" color="dark" label="Sign in" no-caps rounded size="md" style="border-radius: 8px;" type="submit"></q-btn>
        </q-card-section>
      </q-form>
      <!-- Form end -->

      <q-card-section class="text-center q-pt-none">
        <div class="text-grey-8">Don't have an account yet?
          <a class="text-dark text-weight-bold" style="text-decoration: none; cursor: pointer" @click="goToRegister">Sign up.</a>
        </div>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<style scoped>
.my_card {
  width: 25rem;
  border-radius: 8px;
  box-shadow: 0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1);
}
</style>
