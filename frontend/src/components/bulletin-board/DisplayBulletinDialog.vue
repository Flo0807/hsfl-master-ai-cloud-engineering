<template>
  <q-dialog v-model="visible">
    <q-card v-if="props.item" style="width: 800px; max-width: 80vw;">
      <q-card-section class="row items-center q-pb-none">
        <div class="text-h6">{{ item?.title }}</div>
        <q-space />
        <q-btn icon="close" flat round dense v-close-popup />
      </q-card-section>

      <q-card-section>
        <span v-html="item?.content"/>
      </q-card-section>

      <!-- Comment Section -->
      <q-card-section>
        <q-input class="q-pb-md" label="Leave a comment"></q-input>
      </q-card-section>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import {BulletinBoardEntry} from 'components/models';
import {computed, PropType} from 'vue';

const props = defineProps({
  item: {
    type: Object as PropType<BulletinBoardEntry | undefined>,
    default: undefined,
  },
});

const emit = defineEmits(['close']);

const visible = computed({
  get() {
    return props.item !== undefined;
  },
  set() {
    emit('close');
  },
});
</script>
