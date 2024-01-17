<template>
  <q-dialog v-model="dialogModel">
    <q-card style="width: 800px; max-width: 80vw;">
      <q-card-section class="row items-center q-pb-none">
        <div class="text-h6">Edit post #{{form.id}}</div>
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
        <q-btn label="Update" color="primary" @click="update" />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import {computed, PropType, reactive, watch} from 'vue';
import {BulletinBoardEntry} from 'components/models';

const props = defineProps({
  visible: {
    type: Boolean,
    required: true,
    default: false,
  },
  item: {
    type: Object as PropType<BulletinBoardEntry | undefined>,
    default: undefined,
  },
});

const emit = defineEmits(['update', 'close']);

const form = reactive({
  id: '',
  title: '',
  content: '',
} as BulletinBoardEntry);

watch(() => props.item, (item) => {
  if (item) {
    Object.assign(form, item);
  }
});

const dialogModel = computed({
  get() {
    return props.visible;
  },
  set() {
    emit('close');
  },
});

const update = () => {
  emit('update', form);
};

const close = () => {
  emit('close');
};
</script>
