import defaultSettings from '@/settings'

const title = defaultSettings.title || '数据库测试系统'

export default function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${title}`
  }
  return `${title}`
}
