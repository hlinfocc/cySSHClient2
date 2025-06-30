<template>
  <div class="container-list">
    <Breadcrumb :items="['menu.hosts', 'menu.hosts.list']"/>
    <div class="contain">
      <!-- <div class="contain-head">
        <span>{{ $t( 'hosts.form.create' ) }}</span>
        <hr/>
      </div> -->
      <tiny-form
          :model="filterOptions"
          label-position="right"
          label-width="100px"
          class="filter-form"
          size="small"
      >
        <tiny-row :flex="true" justify="center" class="col">
          <tiny-col :span="4" label-width="100px">
            <tiny-form-item :label="$t('hosts.form.name')">
              <tiny-input
                  v-model="filterOptions.hostip"
                  :placeholder="$t('hosts.form.input')"
              ></tiny-input>
            </tiny-form-item>
          </tiny-col>
          <tiny-col :span="4" label-width="100px">
            <tiny-form-item :label="$t('hosts.form.description')">
              <tiny-input
                  v-model="filterOptions.description"
                  :placeholder="$t('hosts.form.input')"
              ></tiny-input>
            </tiny-form-item>
          </tiny-col>
          <tiny-col :span="4">
            <tiny-form-item>
              <div class="search-btn">
                <tiny-button type="info" @click="reloadGrid">
                  {{ $t( 'hosts.form.search' ) }}
                </tiny-button>
                <tiny-button @click="handleFormReset">
                  {{ $t( 'hosts.form.reset' ) }}
                </tiny-button>
              </div>
            </tiny-form-item>
          </tiny-col>
        </tiny-row>

      </tiny-form>
      <div class="bottom-line">
        <hr/>
      </div>
      <tiny-fullscreen
          :teleport="true"
          :page-only="true"
          :z-index="999"
          :fullscreen="fullscreen"
          @update:fullscreen="fullscreen = $event"
      >
        <div class="tiny-fullscreen-scroll">
          <div class="tiny-fullscreen-wrapper">
            <tiny-grid
                ref="taskGrid"
                :fetch-data="fetchDataOption"
                :pager="pagerConfig"
                :loading="loading"
                seq-serial
                size="medium"
                :auto-resize="true"
                row-id="id"
            >
              <template #toolbar>
                <tiny-grid-toolbar>
                  <template #buttons>
                    <div class="btn">
                      <tiny-button type="info"  @click="()=>{editorVisible= true;}">
                          {{ $t('hosts.operation.create') }}
                        </tiny-button>
                      <div class="screen">
                        <img
                            v-if="!fullscreen"
                            src="@/assets/images/screen-out.png"
                            class="screen-image"
                            @click="toggle"
                        />
                        <img
                            v-if="fullscreen"
                            src="@/assets/images/screen-in.png"
                            class="screen-image"
                            @click="toggle"
                        />
                        <span @click="toggle">
                          {{
                            fullscreen
                                ? $t( 'hosts.collapse.restores' )
                                : $t( 'hosts.collapse.full' )
                          }}
                        </span>
                      </div>
                    </div>
                  </template>
                </tiny-grid-toolbar>
              </template>
              <tiny-grid-column type="index" width="60"></tiny-grid-column>
              <tiny-grid-column
                  field="host"
                  :title="$t('hosts.columns.host')"
                  align="left"
              ></tiny-grid-column>
              <tiny-grid-column
                  field="port"
                  :title="$t('hosts.columns.port')"
                  align="center"
              >
            </tiny-grid-column>
            <tiny-grid-column
                  field="username"
                  :title="$t('hosts.columns.username')"
                  align="left"
              ></tiny-grid-column>
              <tiny-grid-column
                  field="iskey"
                  :title="$t('hosts.columns.iskey')"
                  align="center"
              >
                <template #default="{ row }">
                    <tiny-tag type="info"  v-if="row.iskey === 1">是</tiny-tag>
                    <tiny-tag type="warning"  v-if="row.iskey === 0">否</tiny-tag>
                </template>
              </tiny-grid-column>
              <tiny-grid-column
                  field="keyname"
                  :title="$t('hosts.columns.keyname')"
                  align="center"
              ></tiny-grid-column>
              <tiny-grid-column
                  field="hostdesc"
                  :title="$t('hosts.columns.hostdesc')"
                  align="center"
              ></tiny-grid-column>
              <tiny-grid-column
                  :title="$t('hosts.columns.operations')"
                  align="center"
              >
                <template  v-slot="data">
                  <a class="operation-item" @click="handleEditor(data.row)">
                    {{ $t('hosts.columns.operations.editor') }}
                  </a>
                  <a class="operation-item" @click="handleDelete(data.row.id)">
                    {{ $t('hosts.columns.operations.delete') }}
                  </a>
                </template>
              </tiny-grid-column>
            </tiny-grid>
          </div>
        </div>
      </tiny-fullscreen>
    </div>
    <editor v-model:visible="editorVisible" :artdata="artdata" @success="reloadGrid" />
  </div>
</template>

<script lang="ts" setup>
import {ref, reactive, toRefs, onMounted, toRaw} from 'vue';
import {
  Grid as TinyGrid,
  GridColumn as TinyGridColumn,
  GridToolbar as TinyGridToolbar,
  Form as TinyForm,
  FormItem as TinyFormItem,
  Input as TinyInput,
  Button as TinyButton,
  Row as TinyRow,
  Col as TinyCol,
  Select as TinySelect,
  Pager as TinyPager,
  Fullscreen as TinyFullscreen, Modal,TinyTag
} from '@opentiny/vue';
import {
  queryHostsList,
  deleteHosts,
  QueryParmas,
} from '@/api/hosts';

import Editor from './editor.vue';

const editorVisible = ref(false);
const artdata = ref<any>({});

// 初始化请求数据
interface FilterOptions {
  id: string;
  department: string;
  roles: string;
  dateRange: Array<string | Date>;
  name: string;
  status: string;
  workbenchName: string;
  project: string;
  hostip: string;
  description: string;
}

// 加载效果
const state = reactive<{
  loading: boolean;
  filterOptions: FilterOptions;
}>({
  loading: false,
  filterOptions: {} as FilterOptions,
});

const pagerConfig = reactive({
  component: TinyPager,
  attrs: {
    currentPage: 1,
    pageSize: 10,
    pageSizes: [10, 20],
    total: 10,
    layout: 'total, prev, pager, next, jumper, sizes',
  },
});

let tableData = ref([]);
const taskGrid = ref();
const {loading, filterOptions} = toRefs(state);

const statusOptions = [
  {
    value: '0',
    label: '禁用',
  },
  {
    value: '1',
    label: '启用',
  },
];

// 请求数据接口方法
async function fetchData(
    params: QueryParmas = {
      page: 1,
      limit: 10,
      status: '',
    }
) {
  console.log("filterOptions.value:",filterOptions.value);
  
  const {...rest} = filterOptions.value;
  const queryParmas = {
    ...params,
    ...rest,
  };

  state.loading = true;
  try {
    
    const res:any = await queryHostsList(queryParmas);
    
    tableData.value = res.data;
    return {
      result: res.data,
      page: { total: res.count }
    };
  } finally {
    state.loading = false;
  }
}

const fetchDataOption = reactive({
  api: ({page}: any) => {
    const {currentPage, pageSize} = page;

    return fetchData({
      page: currentPage,
      limit:pageSize,
    });
  },
});
const handleDelete = (id:string)=>{
  Modal.confirm('您确定要删除主机吗？不可恢复哦').then((rs:any) => {
    deleteHosts(id).then((res:any)=>{
      if(res.code===200){
        Modal.message({
          message: res.msg,
          status: 'success',
        });
      }else{
        Modal.message({
          message: res.msg,
          status: 'error',
        });
      }
      
    })
  })
  
}
const handleEditor = (e:any)=>{
  editorVisible.value = true;
  console.log("eeeee:",toRaw(e))
  artdata.value = e;
}
function getStatusText(status: string) {
  return statusOptions.find(({value}) => status === value)?.label || '';
}

// form的button
function reloadGrid() {
  taskGrid?.value.handleFetch('reload');
  fetchData();
}

function handleFormReset() {
  state.filterOptions = {} as FilterOptions;
  reloadGrid();
}

const setCollapse = ref(true);

function collapse() {
  setCollapse.value = false;
}

function extend() {
  setCollapse.value = true;
}

// 导出
const toCsvEvent = () => {
  taskGrid.value.exportCsv({
    filename: 'table.csv',
    original: true,
    isHeader: false,
    data: tableData.value,
  });
};

// 全屏缩放设置
const fullscreen = ref(false);
const toggle = () => {
  fullscreen.value = !fullscreen.value;
};
onMounted(()=>{
  fetchData();
})
</script>

<style scoped lang="less">
@import './search-table.less';
</style>
