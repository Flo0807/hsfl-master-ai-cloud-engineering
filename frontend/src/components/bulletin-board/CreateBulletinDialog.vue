<template>
  <q-dialog v-model="dialogModel">
    <q-card style="width: 800px; max-width: 80vw;">
      <q-card-section class="row items-center q-pb-none">
        <div class="text-h6">Create new</div>
        <q-space />
        <q-btn icon="close" flat round dense v-close-popup />
      </q-card-section>

      <q-card-section>
        <q-input class="q-pb-md" label="Titel" v-model="form.title"></q-input>
        <q-editor
          v-model="form.content"
        />
      </q-card-section>
      <q-card-actions class="justify-between">
        <q-btn label="Cancel" flat color="primary" v-close-popup @click="close" />
        <q-btn label="Create" color="primary" @click="create" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import {computed, reactive} from 'vue';
import {BulletinBoardEntry} from 'components/models';

const props = defineProps({
  visible: {
    type: Boolean,
    required: true,
    default: false,
  },
});

const emit = defineEmits(['create', 'close']);

const form = reactive({
  title: '',
  content: '',
} as BulletinBoardEntry);

const dialogModel = computed({
  get() {
    return props.visible;
  },
  set() {
    emit('close');
  },
});

const create = () => {
  emit('create', form);
};

const close = () => {
  emit('close');
};
</script>
