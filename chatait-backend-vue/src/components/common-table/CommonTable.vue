<!--
  - Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
  - Use of this source code is governed by a AGPL v3.0 style
  - license that can be found in the LICENSE file.
  -->

<template>
  <div class="list-common-table">
    <div v-if="title !== ''" class="title-container">
      <h3>{{ title }}</h3>
    </div>
    <div v-if="enableSearch" class="search-container">
      <t-form :data="filterForm" @submit="handleFilterSubmit">
        <div class="search-primary">
          <div
            v-if="tableData !== null && primaryFilterType.field !== '' && filterMode === 'simple'"
            class="search-primary-input"
          >
            <t-form-item :label="primaryFilterType.field_name" :name="primaryFilterType.field" label-align="left">
              <template v-if="primaryFilterType.attr === 'date'">
                <t-date-range-picker
                  v-model="filterForm[primaryFilterType.field].value"
                  clearable
                ></t-date-range-picker>
              </template>
              <template v-else>
                <t-input v-model="filterForm[primaryFilterType.field].value[0]" clearable></t-input>
              </template>
              <t-button
                class="search-primary-input-button-search"
                theme="default"
                variant="base"
                @click="handleFilterSubmit"
              >
                <template #icon>
                  <search-icon />
                </template>
              </t-button>
              <t-button class="search-primary-input-button-advanced" variant="text" @click="handleChangeFilterMode">
                <template v-if="filterMode === 'simple'" #icon>
                  <chevron-down-icon />
                </template>
                <template v-else #icon>
                  <chevron-up-icon />
                </template>
                {{ filterMode === 'simple' ? '高级筛选' : '简单筛选' }}
              </t-button>
            </t-form-item>
          </div>
          <div class="search-primary-tool">
            <t-space :size="0">
              <t-button v-if="showAddButton" variant="outline" @click="handleAdd">
                <template #icon>
                  <add-icon />
                </template>
                {{ addButtonText }}
              </t-button>
              <t-dropdown v-if="rowSelect.enable" :min-column-width="88" trigger="click">
                <t-button variant="outline">
                  选中项
                  <template #icon>
                    <chevron-down-icon></chevron-down-icon>
                  </template>
                </t-button>
                <t-dropdown-menu>
                  <slot name="selected" :selected-rows="selectedRowKeys"></slot>
                </t-dropdown-menu>
              </t-dropdown>
              <t-button v-if="showExportButton" theme="success" variant="outline" @click="handleExport">
                <template #icon>
                  <file-excel-icon />
                </template>
                {{ exportButtonText }}
              </t-button>
            </t-space>
          </div>
        </div>
        <div v-if="tableData !== null && filterMode === 'advanced'" class="search-advanced">
          <t-row :gutter="16">
            <t-col v-for="(item, index) in tableData.filter_types" :key="index" :xs="12" :sm="6">
              <t-form-item :label="item.field_name" :name="item.field" label-align="left">
                <template
                  v-if="
                    filterForm[item.field].model !== undefined &&
                    filterForm[item.field].model !== null &&
                    filterForm[item.field].model === 'date'
                  "
                >
                  <t-date-range-picker v-model="filterForm[item.field].value" class="advanced-input" clearable />
                </template>
                <template
                  v-else-if="
                    filterForm[item.field].model !== undefined &&
                    filterForm[item.field].model !== null &&
                    (filterForm[item.field].model === 'amount' || filterForm[item.field].model === 'money')
                  "
                >
                  <t-range-input v-model="filterForm[item.field].value" class="advanced-input" clearable />
                </template>
                <template
                  v-else-if="
                    filterForm[item.field].model !== undefined &&
                    filterForm[item.field].model !== null &&
                    filterForm[item.field].model === 'select'
                  "
                >
                  <t-select v-model="filterForm[item.field].value[0]" class="advanced-input" clearable>
                    <t-option
                      v-for="(filterFormItem, filterFormIndex) in filterForm[item.field].attrData"
                      :key="filterFormIndex"
                      :label="filterFormItem.label"
                      :value="filterFormItem.value"
                    ></t-option>
                  </t-select>
                </template>
                <template v-else>
                  <t-select v-model="filterForm[item.field].symbol" class="advanced-select" input-value="eq">
                    <t-option key="" label="" value="" />
                    <t-option key="=" label="等于" value="=" />
                    <t-option key="LIKE" label="包含" value="LIKE" />
                    <t-option key="NOTLIKE" label="不包含" value="NOTLIKE" />
                    <t-option key=">" label="大于" value=">" />
                    <t-option key=">=" label="大于等于" value=">=" />
                    <t-option key="<" label="小于" value="<" />
                    <t-option key="<=" label="小于等于" value="<=" />
                    <t-option key="<>" label="不等于" value="<>" />
                  </t-select>
                  <t-input
                    v-model="filterForm[item.field].value[0]"
                    class="advanced-input input-text"
                    clearable
                  ></t-input>
                </template>
              </t-form-item>
            </t-col>
          </t-row>
          <t-row v-if="tableData !== null" :gutter="16">
            <t-col :xs="12" :sm="4">
              <t-button block size="medium" type="submit"> 确定</t-button>
            </t-col>
            <t-col :xs="12" :sm="4">
              <t-button block size="medium" variant="outline" @click="handleClearFilterData"> 重置</t-button>
            </t-col>
            <t-col :xs="12" :sm="4">
              <t-button block size="medium" variant="outline" @click="handleChangeFilterMode"> 简单筛选</t-button>
            </t-col>
          </t-row>
        </div>
      </t-form>
    </div>
    <div class="table-container">
      <t-table
        :data="data"
        :columns="column"
        :row-key="rowKey"
        :vertical-align="verticalAlign"
        :hover="hover"
        :pagination="enablePagination ? pagination : null"
        :loading="dataLoading"
        :size="'small'"
        :selected-row-keys="selectedRowKeys"
        :select-on-row-click="selectOnRowClick"
        :drag-sort="props.dragRow ? 'row' : 'row-handler'"
        :expand-on-row-click="enableExpand"
        @drag-sort="handleSort"
        @select-change="handleSelectChange"
        @page-change="handlePageChange"
        @change="handleTableChange"
      >
        <template v-for="(item, index) in column" :key="index" #[item.colKey]="slotProps">
          <template v-if="!slotsObj[item.colKey]">
            <span>{{ slotProps.row[item.colKey].value }}</span>
          </template>
          <slot :name="item.colKey" :params="slotProps"></slot>
        </template>
        <template v-if="enableExpand" #expanded-row="slotProps">
          <slot name="expandedRow" :params="slotProps"></slot>
        </template>
      </t-table>
    </div>
  </div>
</template>
<script setup lang="ts">
import { PageInfo, PrimaryTableCol, TableRowData } from 'tdesign-vue-next'
import { computed, onMounted, PropType, ref, toRefs, useSlots } from 'vue'
import { AddIcon, ChevronDownIcon, ChevronUpIcon, FileExcelIcon, SearchIcon } from 'tdesign-icons-vue-next'
import { useRoute, useRouter } from 'vue-router'
import { ResponseTableData, ResponseTableDataListValue } from '@/utils/model/response/tableData'
import http from '@/utils/network/http'
import { FilterFormValue, TableFixedColumn, TableRowSelect } from '@/components/common-table/model/model'

// const store = useSettingStore()
const slotsObj = useSlots()
const route = useRoute()
const router = useRouter()

const props = defineProps({
  title: {
    type: String,
    default: '',
  },
  requestPath: {
    type: String,
    default: '',
  },
  rowKey: {
    type: String,
    default: 'id',
  },
  operationColumn: {
    type: Boolean,
    default: false,
  },
  defaultPage: {
    type: Number,
    default: 1,
  },
  defaultPageSize: {
    type: Number,
    default: 20,
  },
  primaryFilterField: {
    type: String,
    default: '',
  },
  showAddButton: {
    type: Boolean,
    default: false,
  },
  addButtonText: {
    type: String,
    default: '添加',
  },
  showExportButton: {
    type: Boolean,
    default: false,
  },
  exportButtonText: {
    type: String,
    default: '导出',
  },
  fixedColumns: {
    type: Array as PropType<TableFixedColumn[]>,
    default: () => [],
  },
  rowSelect: {
    type: Object as PropType<TableRowSelect>,
    default: () => {
      return {
        enable: false,
        type: 'single',
      }
    },
  },
  selectOnRowClick: {
    type: Boolean,
    default: false,
  },
  dragRow: {
    type: Boolean,
    default: false,
  },
  dragRequestPath: {
    type: String,
    default: '',
  },
  enableSearch: {
    type: Boolean,
    default: true,
  },
  enablePagination: {
    type: Boolean,
    default: true,
  },
  enableExpand: {
    type: Boolean,
    default: false,
  },
})

const { rowKey } = toRefs(props)
const verticalAlign = 'middle' as const
const hover = true
const tableData = ref<ResponseTableData | null>(null)

const column = computed((): PrimaryTableCol[] => {
  const reColumn: PrimaryTableCol[] = []
  if (tableData.value !== null && tableData.value.columns.length > 0) {
    for (let i = 0; i < tableData.value.columns.length; i++) {
      reColumn.push({
        title: tableData.value.columns[i].field_name,
        fixed:
          tableData.value.columns[i].field_attr !== null && tableData.value.columns[i].field_attr.fixed !== undefined
            ? tableData.value.columns[i].field_attr.fixed
            : '',
        width:
          tableData.value.columns[i].field_attr !== null && tableData.value.columns[i].field_attr.width !== undefined
            ? tableData.value.columns[i].field_attr.width
            : '',
        ellipsis:
          tableData.value.columns[i].field_attr !== null && tableData.value.columns[i].field_attr.ellipsis !== undefined
            ? tableData.value.columns[i].field_attr.ellipsis
            : false,
        align:
          tableData.value.columns[i].field_attr !== null && tableData.value.columns[i].field_attr.align !== undefined
            ? tableData.value.columns[i].field_attr.align
            : 'left',
        colKey: tableData.value.columns[i].field,
      })
    }
    if (props.dragRow) {
      // reColumn.unshift({
      //   colKey: 'drag', // 列拖拽排序必要参数
      //   title: '排序',
      //   cell: (h) => {
      //     return ''
      //   },
      //   width: 46,
      // })
    }
    if (props.rowSelect.enable) {
      reColumn.unshift({
        colKey: 'row-select',
        type: props.rowSelect.type,
        width: 50,
        fixed: 'left',
      })
    }
    if (props.operationColumn) {
      reColumn.push({
        title: '操作',
        width: 150,
        align: 'center',
        colKey: 'operation',
      })
    }
    // 固定列
    for (let i = 0; i < reColumn.length; i++) {
      for (let j = 0; j < props.fixedColumns.length; j++) {
        if (reColumn[i].colKey === props.fixedColumns[j].field) {
          reColumn[i].fixed = props.fixedColumns[j].direction
          break
        }
      }
    }
  }
  return reColumn
})

const data = computed((): Record<string, ResponseTableDataListValue | string>[] => {
  const reData = []
  if (tableData.value !== null && tableData.value.list.length > 0) {
    for (let i = 0; i < tableData.value.list.length; i++) {
      const tempItem: Record<string, ResponseTableDataListValue> = {}
      for (const itemKey in tableData.value.list[i]) {
        tempItem[itemKey] = tableData.value.list[i][itemKey]
      }
      // if (props.rowSelect.enable && rowKey.value !== '' && tempItem.id !== undefined) {
      //   tempItem[rowKey.value] = tempItem.id.ori_value
      // }
      reData.push(tempItem)
    }
  }
  return reData
})

const pagination = ref({
  current: props.defaultPage,
  pageSize: props.defaultPageSize,
  defaultCurrent: props.defaultPage,
  defaultPageSize: props.defaultPageSize,
  total: 0,
})

const dataLoading = ref(false)

const filterMode = ref('simple') // simple | advanced
const primaryFilterType = ref({
  field_name: '',
  field: '',
  attr: '',
  attr_data: <any>null,
})
const filterForm = ref<Record<string, FilterFormValue>>({})

const queryUrl = () => {
  const urlQuery = { ...route.query }
  if (urlQuery.page_size === undefined) {
    urlQuery.page_size = pagination.value.pageSize.toString()
  }
  // 处理筛选项
  for (const key in urlQuery) {
    if (key === 'page') {
      pagination.value.current = Number.parseInt(urlQuery.page.toString(), 10)
    } else if (key === 'page_size') {
      pagination.value.pageSize = Number.parseInt(urlQuery.page_size.toString(), 10)
    } else if (key === 'filter_mode') {
      filterMode.value = urlQuery.filter_mode.toString()
    } else if (typeof urlQuery[key] === 'string') {
      urlQuery[key] = decodeURIComponent(urlQuery[key].toString())
    }
  }
  return urlQuery
}

const fetchData = async () => {
  dataLoading.value = true
  const requestQuery = queryUrl()
  tableData.value = await http.get(props.requestPath, requestQuery)
  pagination.value.total = tableData.value.total_count
  pagination.value.pageSize = tableData.value.page_size
  // 初始化筛选字段到filterForm
  for (let i = 0; i < tableData.value.filter_types.length; i++) {
    filterForm.value[tableData.value.filter_types[i].field] = {
      symbol: '',
      value: [''],
      model: tableData.value.filter_types[i].attr,
      attrData: tableData.value.filter_types[i].attr_data,
    }
    if (tableData.value.filter_types[i].field === props.primaryFilterField) {
      primaryFilterType.value = tableData.value.filter_types[i]
      filterForm.value[tableData.value.filter_types[i].field].symbol = 'LIKE'
    }
    if (tableData.value.filter_types[i].attr === 'select') {
      filterForm.value[tableData.value.filter_types[i].field].symbol = '='
    }
    if (requestQuery[tableData.value.filter_types[i].field] !== undefined) {
      const tempQueryParamArr = requestQuery[tableData.value.filter_types[i].field].toString().split('|')
      let symbol = ''
      const value = []
      // const model = ''
      for (let j = 0; j < tempQueryParamArr.length; j++) {
        const tempQueryParam = tempQueryParamArr[j].toString().split(',')
        if (tempQueryParam.length > 1) {
          // eslint-disable-next-line prefer-destructuring
          symbol = tempQueryParam[0]
          value.push(tempQueryParam[1])
        }
        if (tempQueryParam.length > 2) {
          // eslint-disable-next-line @typescript-eslint/no-unused-vars
          // model = tempQueryParam[2]
        }
      }
      filterForm.value[tableData.value.filter_types[i].field].symbol = symbol
      filterForm.value[tableData.value.filter_types[i].field].value = value
      // filterForm.value[tableData.value.filter_types[i].field].model = model
    }
  }
  dataLoading.value = false
}

onMounted(() => {
  fetchData()
})

const handlePageChange = async (pageInfo: PageInfo, newDataSource: TableRowData[]) => {
  console.log('分页变化', pageInfo, newDataSource)
  if (pageInfo.pageSize !== pagination.value.pageSize) {
    pagination.value.pageSize = pageInfo.pageSize
    pagination.value.current = 1
    const routeQuery = { ...route.query }
    routeQuery.page = '1'
    routeQuery.page_size = pageInfo.pageSize.toString()
    await router.push({ query: routeQuery })
  } else if (pagination.value.current !== pageInfo.current) {
    pagination.value.current = pageInfo.current
    const routeQuery = { ...route.query }
    routeQuery.page = pageInfo.current.toString()
    await router.push({ query: routeQuery })
  }
  await fetchData()
}
const handleTableChange = () => {
  // console.log('统一Change', changeParams, triggerAndData)
}

const handleRefresh = () => {
  fetchData()
}

const handleChangeFilterMode = async () => {
  if (filterMode.value === 'simple') {
    filterMode.value = 'advanced'
  } else {
    filterMode.value = 'simple'
  }
  const routeQuery = { ...route.query }
  routeQuery.filter_mode = filterMode.value
  await router.push({ query: routeQuery })
}

// 构建筛选请求数据
const handleBuildFilterRequestData = (): Record<string, string> => {
  const requestQuery: Record<string, string> = {}
  for (const key in filterForm.value) {
    if (key === 'level_id') {
      console.log(filterForm.value[key].symbol)
      console.log(filterForm.value[key].value)
      console.log(
        (filterForm.value[key].value.length > 0 && filterForm.value[key].value[0] !== '') ||
          filterForm.value[key].value.length > 1,
      )
    }
    if (
      (filterForm.value[key].value.length > 0 && filterForm.value[key].value[0] !== '') ||
      filterForm.value[key].value.length > 1
    ) {
      requestQuery[key] = ''
      for (let i = 0; i < filterForm.value[key].value.length; i++) {
        if (i !== 0) {
          requestQuery[key] += '|'
        }
        if (
          (filterForm.value[key].model === undefined || filterForm.value[key].model === '') &&
          filterForm.value[key].symbol !== '' &&
          filterForm.value[key].value[i] !== ''
        ) {
          requestQuery[key] += `${filterForm.value[key].symbol},${filterForm.value[key].value}`
          if (filterForm.value[key].model !== undefined && filterForm.value[key].model !== '') {
            requestQuery[key] += `,${filterForm.value[key].model}`
          }
        } else if (filterForm.value[key].model === 'select') {
          requestQuery[key] += `=,${filterForm.value[key].value}`
        } else if (
          filterForm.value[key].model === 'date' ||
          filterForm.value[key].model === 'amount' ||
          filterForm.value[key].model === 'money'
        ) {
          if (i === 0) {
            requestQuery[key] += `>=,${filterForm.value[key].value[0]}`
          } else {
            requestQuery[key] += `<=,${filterForm.value[key].value[1]}`
          }
          if (filterForm.value[key].model !== undefined && filterForm.value[key].model !== '') {
            requestQuery[key] += `,${filterForm.value[key].model}`
          }
        }
      }
    }
  }
  return requestQuery
}

const handleFilterSubmit = async () => {
  // 构建筛选内容
  const requestQuery = handleBuildFilterRequestData()
  const routeQuery = { ...route.query }
  for (const key in routeQuery) {
    if (key !== 'page' && key !== 'page_size' && key !== 'filter_mode') {
      delete routeQuery[key]
    }
  }
  for (const key in requestQuery) {
    if (filterMode.value === 'simple') {
      if (key === props.primaryFilterField) {
        routeQuery[key] = encodeURIComponent(requestQuery[key])
      }
    } else {
      routeQuery[key] = encodeURIComponent(requestQuery[key])
    }
  }
  routeQuery.page = '1'
  await router.push({ query: routeQuery })
  await fetchData()
}

const handleClearFilterData = async () => {
  const routeQuery = { ...route.query }
  for (const key in routeQuery) {
    if (key !== 'page' && key !== 'page_size' && key !== 'filter_mode') {
      delete routeQuery[key]
    }
  }
  await router.push({ query: routeQuery })
  await fetchData()
}

const emit = defineEmits(['onAdd', 'onExport'])

const handleAdd = () => {
  emit('onAdd')
}

const handleExport = () => {
  const requestQuery = queryUrl()
  emit('onExport', requestQuery)
}

const selectedRowKeys = ref<ResponseTableDataListValue[]>([])
const handleSelectChange = (value: any, data: any) => {
  selectedRowKeys.value = value as ResponseTableDataListValue[]
  console.log(value, data)
  console.log(selectedRowKeys.value)
}

const handleSort = async (event: any) => {
  const requestData = {
    sort: <string[]>[],
    page: pagination.value.current,
    page_size: pagination.value.pageSize,
  }
  for (let i = 0; i < event.newData.length; i++) {
    requestData.sort.push(event.newData[i].id.ori_value)
  }
  await http.post(props.dragRequestPath, requestData)
  await fetchData()
}

defineExpose({
  handleRefresh,
})
</script>

<style lang="less">
.list-common-table .search-primary-input .t-input {
  border-radius: 0 !important;
  border-bottom-left-radius: var(--td-radius-default) !important;
  border-top-left-radius: var(--td-radius-default) !important;
}

.list-common-table .search-primary-input .t-range-input {
  border-radius: 0 !important;
  border-bottom-left-radius: var(--td-radius-default) !important;
  border-top-left-radius: var(--td-radius-default) !important;
}

.list-common-table .search-primary-input .t-button {
  border-radius: 0 !important;
  border-bottom-right-radius: var(--td-radius-default) !important;
  border-top-right-radius: var(--td-radius-default) !important;
}

.list-common-table .search-primary-input .search-primary-input-button-search {
  flex: 0 0 auto !important;
  width: 50px !important;
}

.list-common-table .search-primary-input .search-primary-input-button-advanced {
  flex: 0 0 auto !important;
  width: 90px !important;
  padding-left: 0 !important;
  padding-right: 0 !important;
}

.list-common-table .search-primary-tool .t-space .t-button {
  border-radius: 0 !important;
  border-right: none;
}

.list-common-table .search-primary-tool .t-space .t-space-item:first-child .t-button {
  border-top-left-radius: var(--td-radius-default) !important;
  border-bottom-left-radius: var(--td-radius-default) !important;
}

.list-common-table .search-primary-tool .t-space .t-space-item:last-child .t-button {
  border-top-right-radius: var(--td-radius-default) !important;
  border-bottom-right-radius: var(--td-radius-default) !important;
  border-right: 1px solid;
}

.list-common-table .search-advanced .t-col {
  margin-top: 10px;
}

.list-common-table .search-advanced .advanced-select .t-input {
  border-radius: 0 !important;
  border-top-left-radius: var(--td-radius-default) !important;
  border-bottom-left-radius: var(--td-radius-default) !important;
  border-right: none !important;
}

.list-common-table .search-advanced .advanced-input.input-text .t-input {
  border-radius: 0 !important;
  border-top-right-radius: var(--td-radius-default) !important;
  border-bottom-right-radius: var(--td-radius-default) !important;
}
</style>

<style lang="less" scoped>
.list-common-table {
  background-color: var(--td-bg-color-container);
  padding: var(--td-comp-paddingTB-xxl) var(--td-comp-paddingLR-xxl);
  border-radius: var(--td-radius-medium);

  .table-container {
    margin-top: var(--td-comp-margin-xxl);
  }

  .search-container {
    width: 100%;

    .search-primary {
      width: 100%;
      overflow: hidden;

      .search-primary-input {
        width: 100%;
      }

      .search-primary-tool {
        width: 100%;
        display: flex;
        align-items: center;

        overflow: hidden;
      }
    }

    @media screen and (min-width: 768px) {
      .search-primary {
        display: flex;

        .search-primary-input {
          flex: 0 0 auto;
          width: 500px;
        }

        .search-primary-tool {
          flex: 1 1 auto;
          width: 100%;
          justify-content: end;
        }
      }
    }

    .search-advanced {
      padding-bottom: 10px;
      border-bottom: 1px solid var(--td-border-level-2-color);

      .advanced-select {
        flex: 0 0 auto;
        width: 70px;
        overflow: hidden;
      }

      .advanced-input {
        flex: 1 1 auto;
        width: 100%;
      }
    }
  }
}
</style>
