export class StringUtils {
  static titleCase(str: string) {
    return str.charAt(0).toUpperCase() + str.slice(1);
  }
}
