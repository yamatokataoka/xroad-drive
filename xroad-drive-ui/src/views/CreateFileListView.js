import FileListView from '@/views/FileListView'

// This is a factory function for dynamically creating root-level file list
// views, since they share most of the logic except for the type of items to
// display. They are essentially higher order components wrapping ItemList.vue.
export default function createFileListView(type) {
  return {
    name: type,
    render(h) {
      return h(FileListView, {
        props: {
          type
        }
      });
    }
  }
}