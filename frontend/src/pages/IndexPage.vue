<template>
  <q-page class="row items-center justify-evenly">
    <q-infinite-scroll class="q-mt-lg" @load="onLoad" :offset="500" :disable="bulletinStore.reachedEnd || bulletinStore.isLoading">
      <div v-for="(entry, index) in bulletinStore.get" :key="index">
        <q-card
          flat bordered
          @click="displayBulletinBoardEntry = entry"
          style="min-width: 700px"
          class="q-mb-lg"
        >
          <q-card-section horizontal>
            <q-card-section class="q-pt-xs">
              <div class="text-overline">{{ formatCreatedAt(entry.createdAt) }}</div>
              <div class="text-h5 q-mt-sm q-mb-xs">{{ entry.title }}</div>
              <div class="text-caption text-grey" v-html="entry.content"/>
            </q-card-section>
          </q-card-section>

          <q-separator />

          <q-card-actions>
            <q-space/>
            <q-btn-group flat>
              <q-btn icon="edit" @click.stop="() => {
                editBulletinBoardEntry = entry
                displayEditDialog = true
              }" />
              <q-btn text-color="red" icon="delete" @click.stop="confirmDelete(entry)" />
            </q-btn-group>
          </q-card-actions>
        </q-card>
      </div>
      <template v-slot:loading>
        <div class="row justify-center q-my-md">
          <q-spinner-dots color="primary" size="40px" />
        </div>
      </template>
    </q-infinite-scroll>
    <q-page-sticky position="bottom-right" :offset="[18, 18]">
      <q-btn fab icon="add" color="accent" @click="displayCreateDialog = true" />
    </q-page-sticky>

    <!-- Create Dialog -->
    <CreateBulletinDialog :visible="displayCreateDialog" @close="displayCreateDialog = false" @create="(post) => {
      bulletinStore.create(post)
      displayCreateDialog = false
    }" />
    <EditBulletinDialog
      :visible="displayEditDialog"
      :item="editBulletinBoardEntry"
      @close="displayEditDialog = false"
      @update="(post) => {
        bulletinStore.update(post)
        displayEditDialog = false
      }"
    />
    <!-- Display Dialog -->
    <DisplayBulletinDialog :item="displayBulletinBoardEntry" @close="displayBulletinBoardEntry = undefined" />

  </q-page>
</template>

<script setup lang="ts">
import CreateBulletinDialog from 'components/bulletin-board/CreateBulletinDialog.vue';
import {DateTime} from 'luxon';

import {onMounted, ref} from 'vue';
import {BulletinBoardEntry} from 'components/models';
import DisplayBulletinDialog from 'components/bulletin-board/DisplayBulletinDialog.vue';
import {useBulletinStore} from 'stores/bulletin-store';
import EditBulletinDialog from 'components/bulletin-board/EditBulletinDialog.vue';
import {useQuasar} from 'quasar';

const $q = useQuasar()
const bulletinStore = useBulletinStore()

const displayCreateDialog = ref(false)
const displayEditDialog = ref(false)
const displayBulletinBoardEntry = ref<BulletinBoardEntry | undefined>(undefined)
const editBulletinBoardEntry= ref<BulletinBoardEntry | undefined>(undefined)

const formatCreatedAt = (createdAt: string) => {
  return DateTime.fromISO(createdAt).toLocaleString(DateTime.DATETIME_MED)
}

const onLoad = (index: any, done: () => void) => {
  setTimeout(() => {
    bulletinStore.fetch()
    done()
  }, 200)
}

onMounted(() => {
  bulletinStore.fetch()
})


const confirmDelete = (item: BulletinBoardEntry) => {
  $q.dialog({
    title: 'Confirm',
    message: `Would you like to remove post #${item.id}?`,
    cancel: 'Cancel',
    ok: 'Remove',
    persistent: false
  }).onOk(() => {
    bulletinStore.remove(item.id)
  })
}

</script>
