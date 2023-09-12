# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: BUSL-1.1

sample "build_ce_linux_amd64_deb" {
  attributes = global.sample_attributes

  subset "agent" {
    matrix {
      arch            = ["amd64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["ubuntu"]
      edition         = ["ce"]
    }
  }

  subset "smoke" {
    matrix {
      arch            = ["amd64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["ubuntu"]
      edition         = ["ce"]
    }
  }

  subset "proxy" {
    matrix {
      arch            = ["amd64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["ubuntu"]
      edition         = ["ce"]
    }
  }

  subset "upgrade" {
    matrix {
      arch            = ["amd64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["ubuntu"]
      edition         = ["ce"]
    }
  }
}

sample "build_ce_linux_arm64_deb" {
  attributes = global.sample_attributes

  subset "agent" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["ubuntu"]
      edition         = ["ce"]
    }
  }

  subset "smoke" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["ubuntu"]
      edition         = ["ce"]
    }
  }

  subset "proxy" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["ubuntu"]
      edition         = ["ce"]
    }
  }

  subset "upgrade" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["ubuntu"]
      edition         = ["ce"]
    }
  }
}

sample "build_ce_linux_arm64_rpm" {
  attributes = global.sample_attributes

  subset "agent" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["rhel"]
      edition         = ["ce"]
    }
  }

  subset "smoke" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["rhel"]
      edition         = ["ce"]
    }
  }

  subset "proxy" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["rhel"]
      edition         = ["ce"]
    }
  }

  subset "upgrade" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["rhel"]
      edition         = ["ce"]
    }
  }
}

sample "build_ce_linux_amd64_rpm" {
  attributes = global.sample_attributes

  subset "agent" {
    matrix {
      arch            = ["amd64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["rhel"]
      edition         = ["ce"]
    }
  }

  subset "smoke" {
    matrix {
      arch            = ["amd64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["rhel"]
      edition         = ["ce"]
    }
  }

  subset "proxy" {
    matrix {
      arch            = ["amd64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["rhel"]
      edition         = ["ce"]
    }
  }

  subset "upgrade" {
    matrix {
      arch            = ["amd64"]
      artifact_source = ["crt"]
      artifact_type   = ["package"]
      distro          = ["rhel"]
      edition         = ["ce"]
    }
  }
}

sample "build_ce_linux_amd64_zip" {
  attributes = global.sample_attributes

  subset "agent" {
    matrix {
      arch            = ["amd64"]
      artifact_type   = ["bundle"]
      artifact_source = ["crt"]
      edition         = ["ce"]
    }
  }

  subset "smoke" {
    matrix {
      arch            = ["amd64"]
      artifact_type   = ["bundle"]
      artifact_source = ["crt"]
      edition         = ["ce"]
    }
  }

  subset "proxy" {
    matrix {
      arch            = ["amd64"]
      artifact_type   = ["bundle"]
      artifact_source = ["crt"]
      edition         = ["ce"]
    }
  }

  subset "upgrade" {
    matrix {
      arch            = ["amd64"]
      artifact_type   = ["bundle"]
      artifact_source = ["crt"]
      edition         = ["ce"]
    }
  }
}

sample "build_ce_linux_arm64_zip" {
  attributes = global.sample_attributes

  subset "agent" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["bundle"]
      edition         = ["ce"]
    }
  }

  subset "smoke" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["bundle"]
      edition         = ["ce"]
    }
  }

  subset "proxy" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["bundle"]
      edition         = ["ce"]
    }
  }

  subset "upgrade" {
    matrix {
      arch            = ["arm64"]
      artifact_source = ["crt"]
      artifact_type   = ["bundle"]
      edition         = ["ce"]
    }
  }
}
