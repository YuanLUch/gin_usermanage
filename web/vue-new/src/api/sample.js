import request from '@/utils/request'

export default {
    uploadFolder(caseSet) {
        return request({
            url: '/sample/uploadFolder',
            method: 'post',
            data: caseSet
        })
    },
    uploadFiles(cases) {
        return request({
            url: '/sample/uploadFiles',
            method: 'post',
            data: cases
        })
    },
    updateFiles(cases) {
        return request({
            url: '/sample/updateFiles',
            method: 'put',
            data: cases
        })
    },
    getCaseSetList(searchModel) {
        return request({
            url: '/sample/list',
            method: 'get',
            params:
            {
                pageNo: searchModel.pageNo,
                pageSize: searchModel.pageSize,
                name: searchModel.name,
            }
        })
    },

    getCaseList(searchModel) {
        return request({
            url: '/sample/case-list',
            method: 'get',
            params:
            {
                pageNo: searchModel.pageNo,
                pageSize: searchModel.pageSize,
                case_name: searchModel.caseName,
                case_set_id: searchModel.case_set_id
            }
        })
    },
    deleteSetById(id) {
        return request({
          url: `/set/${id}`,
          method: 'delete'
        });
    },
    deleteCaseById(id) {
        return request({
          url: `/sample/case/${id}`,
          method: 'delete'
        });
      },
    getCaseById(id){
        return request({
            url: `/sample/case-info/${id}`,
            method: `get`
        })
    },
    saveCase(case1) {
        if (case1.id == null && case1.id == undefined) {
          return this.uploadFiles(case1)
        }
        return this.updateFiles(case1)
      }
}
