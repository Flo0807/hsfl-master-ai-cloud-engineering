<script lang="ts" setup>
import {ref} from 'vue';
import {useRouter} from 'vue-router';
import {useAuthStore} from 'stores/auth-store';
import {useQuasar} from 'quasar';

const $q = useQuasar();
const router = useRouter()
const authStore = useAuthStore()

const email = ref('')
const password = ref('')
const passwordRepeated = ref('')

const register = () => {
  if (password.value === passwordRepeated.value) {
    authStore.register(
      {
        email: email.value,
        password: password.value
      }
    ).then(() => {
      resetInputs();
      goToLogin();
    }).catch(() => {
      $q.notify({
        color: 'negative',
        position: 'top',
        message: 'Register failed',
        icon: 'report_problem'
      })
    })
  }
}

const resetInputs = () => {
  email.value = '';
  password.value = '';
  passwordRepeated.value = '';
}

const goToLogin = () => {
  router.push({name: 'login'})
}
</script>

<template>
  <q-page class="flex flex-center bg-grey-2">
    <q-card bordered class="q-pa-md shadow-2 my_card">
      <q-card-section class="text-center">
        <div class="text-grey-9 text-h5 text-weight-bold">Register</div>
        <div class="text-grey-8">Register below to access the application</div>
      </q-card-section>

      <!-- Form start -->
      <q-form @submit.prevent="register">
        <q-card-section>
          <q-input v-model="email" dense label="Email Address" outlined type="email"></q-input>
          <q-input v-model="password" class="q-mt-md" dense label="Password" outlined type="password"></q-input>
          <q-input v-model="passwordRepeated" class="q-mt-md" dense label="Repeat Password" outlined type="password"></q-input>
        </q-card-section>

        <q-card-section>
          <q-btn class="full-width" color="dark" label="Register" no-caps rounded size="md" style="border-radius: 8px;" type="submit"></q-btn>
        </q-card-section>
      </q-form>
      <!-- Form end -->

      <q-card-section class="text-center q-pt-none">
        <div class="text-grey-8">Already have an account?
          <a class="text-dark text-weight-bold" style="text-decoration: none; cursor: pointer" @click="goToLogin">Sign in.</a>
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
