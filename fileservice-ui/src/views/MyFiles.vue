<template>
  <div>
    <tool-bar />
    <file-list />
  </div>
</template>

<script>
import FileList from '@/components/FileList';
import ToolBar from '@/components/ToolBar';

export default {
  name: 'MyFiles',
  components: { FileList, ToolBar },
}
</script>
