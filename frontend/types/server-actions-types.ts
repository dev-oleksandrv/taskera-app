export interface ServerActionResult {
  success: boolean;
  errors?: string[];
  fieldErrors?: Record<string, string[]>;
}