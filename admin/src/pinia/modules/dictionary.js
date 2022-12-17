import { findDictionaryByKey } from '@/api/dictionary'

import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useDictionaryStore = defineStore('dictionary', () => {
  const dictionaryMap = ref({})

  // 单个key查找
  const getDictionary = async (key) => {
    if (dictionaryMap.value[key] && dictionaryMap.value[key].length) {
      return dictionaryMap.value[key]
    } else {
      const res = await findDictionaryByKey(key)
      if (res.code === 0) {
        const dict = []
        res.data && res.data.forEach(item => {
          dict.push({
            label: item.name,
            value: item.id
          })
        })
        dictionaryMap.value[key] = dict
        return dictionaryMap.value[key]
      }
    }
  }

  return {
    getDictionary
  }
})
