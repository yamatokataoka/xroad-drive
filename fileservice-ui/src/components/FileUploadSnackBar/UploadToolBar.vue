<template>
  <v-toolbar>
    <v-btn
      icon
      @click="emitClickClose"
    >
      <v-icon>mdi-close</v-icon>
    </v-btn>
    <v-toolbar-title class="body-2">Uploading...</v-toolbar-title>
    <v-spacer></v-spacer>
    <v-btn
      icon
      v-if="listOpen"
      @click="emitClickChevronUp"
    ><v-icon>mdi-chevron-up</v-icon>
    </v-btn>
    <v-btn
      icon
      v-else
      @click="emitClickChevronDown"
    ><v-icon>mdi-chevron-down</v-icon>
    </v-btn>
  </v-toolbar>
</template>

<script>
  export default {
    name: 'UploadToolBar',
    props: {
      listOpen: {
        required: true
      }
    },
    data() {
      return {
        value: 50,
        uploadFinished: false
      }
    },
    methods: {
      emitClickClose() {
        this.$emit('clickClose');
      },
      emitClickChevronUp() {
        this.$emit('clickChevronUp');
      },
      emitClickChevronDown() {
        this.$emit('clickChevronDown');
      }
    }
  };
</script>