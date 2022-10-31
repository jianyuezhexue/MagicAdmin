import { findDictionary } from '@/api/Dictionary'

import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useDictionaryStore = defineStore('dictionary', () => {
  const dictionaryMap = ref({})

  const setDictionaryMap = (dictionaryRes) => {
    dictionaryMap.value = { ...dictionaryMap.value, ...dictionaryRes }
  }

  const getDictionary = async(type) => {
    if (dictionaryMap.value[type] && dictionaryMap.value[type].length) {
      return dictionaryMap.value[type]
    } else {
      const res = await findDictionary({ type })
      if (res.code === 0) {
        const dictionaryRes = {}
        const dict = []
        res.data.reDictionary.DictionaryDetails && res.data.reDictionary.DictionaryDetails.forEach(item => {
          dict.push({
            label: item.label,
            value: item.value
          })
        })
        dictionaryRes[res.data.reDictionary.type] = dict
        setDictionaryMap(dictionaryRes)
        return dictionaryMap.value[type]
      }
    }
  }

  return {
    dictionaryMap,
    setDictionaryMap,
    getDictionary
  }
})
