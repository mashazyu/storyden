import { Button } from "src/theme/components/Button";

import { useControl } from "./useControl";

export function Bold() {
  const { isActive, onToggle } = useControl("bold");

  return (
    <Button
      backgroundColor={isActive ? "blackAlpha.100" : undefined}
      onMouseDown={onToggle}
      aria-label="Bold"
    >
      <svg
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M13.259 17.625H7.47778V6.375H12.7903C13.4166 6.37504 14.0299 6.55435 14.5576 6.89174C15.0853 7.22914 15.5054 7.71052 15.7683 8.27903C16.0312 8.84754 16.1259 9.47942 16.0412 10.1C15.9565 10.7206 15.6959 11.304 15.2903 11.7812C15.8201 12.205 16.2056 12.7825 16.3936 13.4344C16.5816 14.0862 16.563 14.7803 16.3403 15.4211C16.1175 16.0619 15.7016 16.6179 15.1498 17.0126C14.598 17.4073 13.9374 17.6213 13.259 17.625ZM9.35278 15.75H13.2465C13.4312 15.75 13.6141 15.7136 13.7847 15.643C13.9553 15.5723 14.1103 15.4687 14.2409 15.3381C14.3715 15.2075 14.4751 15.0525 14.5457 14.8819C14.6164 14.7113 14.6528 14.5284 14.6528 14.3438C14.6528 14.1591 14.6164 13.9762 14.5457 13.8056C14.4751 13.635 14.3715 13.48 14.2409 13.3494C14.1103 13.2188 13.9553 13.1152 13.7847 13.0445C13.6141 12.9739 13.4312 12.9375 13.2465 12.9375H9.35278V15.75ZM9.35278 11.0625H12.7903C12.975 11.0625 13.1578 11.0261 13.3284 10.9555C13.499 10.8848 13.6541 10.7812 13.7847 10.6506C13.9152 10.52 14.0188 10.365 14.0895 10.1944C14.1602 10.0238 14.1965 9.84092 14.1965 9.65625C14.1965 9.47158 14.1602 9.28872 14.0895 9.1181C14.0188 8.94749 13.9152 8.79246 13.7847 8.66188C13.6541 8.5313 13.499 8.42772 13.3284 8.35704C13.1578 8.28637 12.975 8.25 12.7903 8.25H9.35278V11.0625Z"
          fill="#212529"
        />
      </svg>
    </Button>
  );
}
