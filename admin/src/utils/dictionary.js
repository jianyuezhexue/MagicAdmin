import { useDictionaryStore } from '@/pinia/modules/dictionary'
//  获取字典方法 使用示例 getDict('sex').then(res)  或者 async函数下 const res = await getDict('sex')
export const getDict = async (key) => {
  const dictionaryStore = useDictionaryStore()
  const res = await dictionaryStore.getDictionary(key)
  return res
}
